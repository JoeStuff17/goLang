package user_s

import (
	// "database/sql"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	// "strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"main.go/database"
	"main.go/enums"
	"main.go/helpers"
	dto "main.go/interface_model"
	"main.go/models"
)

func CreateUser(payload *models.Users) dto.GenericResponse {
	err := database.DBSql.Create(payload).Error
	if err != nil {
		return dto.GenericResponse{Success: false, Data: err}
	}
	return dto.GenericResponse{
		Success: true,
		Data:    payload.ID,
	}
}

func FetchAllUsers() dto.GenericResponse {
	var users []models.Users
	err := database.DBSql.Limit(2).Offset(0).Find(&users).Error
	if err != nil {
		return dto.GenericResponse{
			Success: false,
			Data:    err,
		}
	}
	return dto.GenericResponse{
		Success: true,
		Data:    users,
	}
}

func SendOtp(payload *dto.SendOtpPayload) dto.GenericResponse {
	var user *models.Users
	var admin *models.Admins
	var church_user *models.ChurchUser
	var roles []any
	if strings.TrimSpace(payload.MobileNumber) == "" {
		return dto.GenericResponse{Success: false, Data: nil, Message: "Mobile number is required"}
	}
	payload.MobileNumber = strings.TrimSpace(payload.MobileNumber)
	if payload.Role == "" {
		return dto.GenericResponse{Success: false, Data: nil, Message: "Role is required"}
	}

	switch payload.Role {
	case string(enums.RoleAdmin):
		roles = []any{enums.RoleAdmin, enums.RoleSuperAdmin}
	case string(enums.RoleChurchUser):
		roles = []any{enums.RoleChurchAdmin, enums.RoleChurchUser}
	default:
		return dto.GenericResponse{Success: false, Message: "Invalid role"}
	}

	database.DBSql.Model(&models.Users{}).Select("id, lastOTPSent").Where("mobile_number = ? AND role IN ?", payload.MobileNumber, roles).First(&user)
	// Verify User
	if user == nil {
		return dto.GenericResponse{Success: false, Data: nil, Message: "Not allowed to perform this action", StatusCode: fiber.StatusBadRequest}
	}

	// Check last log-in time to avoid spam attacks
	if user.LastOTPSent.Valid && time.Since(user.LastOTPSent.Time) < 30*time.Second {
		wait := 30 - int(time.Since(user.LastOTPSent.Time).Seconds())
		return dto.GenericResponse{Success: false, Message: fmt.Sprintf("Please try again in %d seconds", wait)}
	}

	// Verify Admin
	if payload.Role == string(enums.RoleAdmin) {
		database.DBSql.Model(&models.Admins{}).Select("id, status").Where("mobile_number = ?", payload.MobileNumber).First(&admin)
		if admin == nil {
			return dto.GenericResponse{Success: false, Data: nil, Message: "Not allowed to perform this action", StatusCode: fiber.StatusBadRequest}
		} else if admin.Status != enums.ACTIVE {
			return dto.GenericResponse{Success: false, Data: nil, Message: fmt.Sprintf("Account is in %s", admin.Status), StatusCode: fiber.StatusBadRequest}
		}
	} else if payload.Role == string(enums.RoleChurchUser) {
		database.DBSql.Model(&models.ChurchUser{}).Select("id, is_active").Where("mobile_number = ?", payload.MobileNumber).First(&church_user)
		if church_user == nil {
			return dto.GenericResponse{Success: false, Data: nil, Message: "Not allowed to perform this action", StatusCode: fiber.StatusBadRequest}
		} else if !church_user.IsActive {
			return dto.GenericResponse{Success: false, Data: nil, Message: "Account is not active", StatusCode: fiber.StatusBadRequest}
		}
	}
	// Generate OTP & Trigger SMS
	otp := helpers.SendLoginOtp(payload.MobileNumber)
	otpNum, err := strconv.Atoi(otp)
	if err != nil {
		fmt.Println("Error:", err)
	}
	database.DBSql.Model(&models.Users{}).Where("id=?", user.ID).Updates(models.Users{OTP: otpNum, LastOTPSent: sql.NullTime{Time: time.Now().UTC(), Valid: true}})
	if os.Getenv("ENV") == "production" {
		otpNum = 0000
	}
	return dto.GenericResponse{
		Success:    true,
		Data:       otpNum,
		Message:    "OTP sent successfully!",
		StatusCode: fiber.StatusOK,
	}
}

func VerifyOtp(payload *dto.VerifyOtpPayload) dto.GenericResponse {
	var user models.Users
	var token, refreshToken string
	var roleSet []any
	isAllowed := true

	// Role-based support
	switch payload.Role {
	case string(enums.RoleAdmin):
		roleSet = []any{enums.RoleAdmin, enums.RoleSuperAdmin}
	case string(enums.RoleChurchUser):
		roleSet = []any{enums.RoleChurchAdmin, enums.RoleChurchUser}
	default:
		return dto.GenericResponse{Success: false, Message: "Invalid role", StatusCode: fiber.StatusBadRequest}
	}

	// Verify User by OTP
	err := database.DBSql.Model(&models.Users{}).
		Select("id, lastOTPSent, name, role").
		Where("mobile_number = ? AND role IN ? AND otp = ?", payload.MobileNumber, roleSet, payload.Otp).
		Take(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.GenericResponse{Success: false, Message: "Invalid OTP or user not found", StatusCode: fiber.StatusBadRequest}
	}

	// Role-specific entity validation
	if user.Role == enums.RoleAdmin || user.Role == enums.RoleSuperAdmin {
		var admin models.Admins
		err := database.DBSql.Model(&models.Admins{}).
			Select("id, name, status, is_device_restricted, is_location_restricted, allowed_devices, allowed_locations").
			Where("mobile_number = ?", payload.MobileNumber).
			Take(&admin).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.GenericResponse{Success: false, Message: "Admin not found", StatusCode: fiber.StatusBadRequest}
		}
		if admin.Status != enums.ACTIVE {
			return dto.GenericResponse{Success: false, Message: fmt.Sprintf("Account is in %s", admin.Status), StatusCode: fiber.StatusBadRequest}
		}

		// Check device restrictions
		if admin.IsDeviceRestricted {
			type Iot struct {
				BrowserName string `json:"browserName"`
				BrowserId   string `json:"browserId"`
			}
			var devices []Iot
			_ = json.Unmarshal([]byte(admin.AllowedDevices.String()), &devices)
			isAllowed = false
			for _, d := range devices {
				if d.BrowserId == payload.BrowserId {
					isAllowed = true
					break
				}
			}
		}

		// Location restrictions (not applied but can be logged)
		if admin.IsLocationRestricted {
			type Iot struct {
				Address string  `json:"address"`
				Lat     float64 `json:"lat"`
				Lng     float64 `json:"lng"`
				Type    string  `json:"type"`
			}
			var location Iot
			_ = json.Unmarshal([]byte(admin.AllowedLocations.String()), &location)
			// Log / extend later
		}

		token = helpers.GenerateToken(strconv.Itoa(int(user.ID)), string(user.Role), user.Name)
		refreshToken = helpers.GenerateRefreshToken(strconv.Itoa(int(user.ID)))
		database.DBSql.Model(&models.Users{}).Where("id = ?", user.ID).Updates(models.Users{Token: &token, RefreshToken: &refreshToken})

		if isAllowed {
			return dto.GenericResponse{
				Success: true,
				Data: fiber.Map{
					"user":  fiber.Map{"id": user.ID, "name": user.Name},
					"admin": fiber.Map{"id": admin.ID, "name": admin.Name},
					"token": token, "refreshToken": refreshToken,
				},
				Message:    "Logged-in successfully!",
				StatusCode: fiber.StatusOK,
			}
		}
		return dto.GenericResponse{Success: false, Message: "Login restricted", StatusCode: fiber.StatusForbidden}
	}

	// Church User validation
	if user.Role == enums.RoleChurchAdmin || user.Role == enums.RoleChurchUser {
		var churchUser models.ChurchUser
		err := database.DBSql.Model(&models.ChurchUser{}).
			Select("id, name, is_active").
			Where("mobile_number = ?", payload.MobileNumber).
			Take(&churchUser).Error

		if errors.Is(err, gorm.ErrRecordNotFound) || !churchUser.IsActive {
			return dto.GenericResponse{Success: false, Message: "Church user not found or inactive", StatusCode: fiber.StatusBadRequest}
		}

		token = helpers.GenerateToken(strconv.Itoa(int(user.ID)), string(user.Role), user.Name)
		refreshToken = helpers.GenerateRefreshToken(strconv.Itoa(int(user.ID)))
		database.DBSql.Model(&models.Users{}).Where("id = ?", user.ID).Updates(models.Users{Token: &token, RefreshToken: &refreshToken})

		return dto.GenericResponse{
			Success: true,
			Data: fiber.Map{
				"user":       fiber.Map{"id": user.ID, "name": user.Name},
				"churchUser": fiber.Map{"id": churchUser.ID, "name": churchUser.Name},
				"token":      token, "refreshToken": refreshToken,
			},
			Message:    "Logged-in successfully!",
			StatusCode: fiber.StatusOK,
		}
	}

	return dto.GenericResponse{Success: false, Message: "Unhandled role", StatusCode: fiber.StatusInternalServerError}
}

func FetchUserById(user_id int) dto.GenericResponse {
	var userDetails dto.FetchUserWithDetails
	query := `SELECT 
		u.id as user_id,
		u.role,
		u.related_id,
		a.name AS admin_name,
		c.name AS member_name,
		c2.id as church_id,
		c2.name as church_name
	FROM users u
	LEFT JOIN admins a ON u.role IN ('admin', 'super_admin') AND u.related_id = a.id
	LEFT JOIN church_users c ON u.role NOT IN ('admin', 'super_admin') AND u.related_id = c.id
	left join churches c2 on c2.id = c.church_id
	WHERE u.id = ?
	LIMIT 1`

	err := database.DBSql.Raw(query, user_id).Scan(&userDetails).Error
	if err != nil {
		return dto.GenericResponse{Success: false, Message: err.Error(), Data: nil, StatusCode: fiber.StatusInternalServerError}
	}

	if userDetails.UserId == 0 {
		return dto.GenericResponse{
			Success: true, Message: "No records found", Data: nil, StatusCode: fiber.StatusOK,
		}
	}

	return dto.GenericResponse{
		Success: true, Message: "User retrieved successfully", Data: userDetails, StatusCode: fiber.StatusOK,
	}
}
