package user_s

import (
	// "database/sql"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	// "strconv"
	"time"

	"github.com/gofiber/fiber/v2"
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

func AdminSendOtp(payload *dto.AdminSendOtpPayload) dto.GenericResponse {
	var user *models.Users
	var admin *models.Admins
	if strings.TrimSpace(payload.MobileNumber) == "" {
		return dto.GenericResponse{Success: false, Data: nil, Message: "Mobile number is required"}
	}
	payload.MobileNumber = strings.TrimSpace(payload.MobileNumber)
	user_role := enums.RoleAdmin
	if payload.MobileNumber == "9962675336" {
		user_role = enums.RoleSuperAdmin
	}
	database.DBSql.Model(&models.Users{}).Select("id, lastOTPSent").Where(&models.Users{MobileNumber: payload.MobileNumber, Role: user_role}).Scan(&user)
	// Verify User
	if user == nil {
		return dto.GenericResponse{Success: false, Data: nil, Message: "Not allowed to perform this action", StatusCode: fiber.StatusBadRequest}
	}
	// Check last login time
	if user.LastOTPSent.Valid {
		timeDiffInLogin := time.Now().Sub(user.LastOTPSent.Time)
		if timeDiffInLogin.Seconds() < 30 {
			tryAgainInSecInFloat := 30 - timeDiffInLogin.Seconds()
			tryAgainInSec := fmt.Sprintf("%.2f", tryAgainInSecInFloat)
			return dto.GenericResponse{Success: false, Data: nil, Message: "Please try again in " + tryAgainInSec + " seconds", StatusCode: fiber.StatusBadRequest}
		}
	}
	// Verify Admin
	database.DBSql.Model(&models.Admins{}).Select("id, status").Where(&models.Admins{MobileNumber: payload.MobileNumber}).Scan(&admin)
	if admin == nil {
		return dto.GenericResponse{Success: false, Data: nil, Message: "Not allowed to perform this action", StatusCode: fiber.StatusBadRequest}
	} else if admin.Status != enums.ACTIVE {
		return dto.GenericResponse{Success: false, Data: nil, Message: fmt.Sprintf("Account is in %s", admin.Status), StatusCode: fiber.StatusBadRequest}
	}
	// Generate OTP & Trigger SMS
	otp := helpers.SendLoginOtp(admin.MobileNumber)
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

func AdminVerifyOtp(payload *dto.AdminVerifyOtpPayload) dto.GenericResponse {
	var user *models.Users
	var admin *models.Admins
	isAllowed := true
	// Verify OTP
	user_role := enums.RoleAdmin
	if payload.MobileNumber == "9962675336" {
		user_role = enums.RoleSuperAdmin
	}
	database.DBSql.Model(&models.Users{}).Select("id, lastOTPSent, name, role").Where(&models.Users{MobileNumber: payload.MobileNumber, Role: user_role, OTP: payload.Otp}).Scan(&user)
	if user == nil {
		return dto.GenericResponse{Success: false, Data: nil, Message: "Not allowed to perform this action", StatusCode: fiber.StatusBadRequest}
	}

	// Verify Admin
	database.DBSql.Model(&models.Admins{}).Select("id, status, name").Where(&models.Admins{MobileNumber: payload.MobileNumber}).Scan(&admin)
	if admin == nil {
		return dto.GenericResponse{Success: false, Data: nil, Message: "Not allowed to perform this action", StatusCode: fiber.StatusBadRequest}
	} else if admin.Status != enums.ACTIVE {
		return dto.GenericResponse{Success: false, Data: nil, Message: fmt.Sprintf("Account is in %s", admin.Status), StatusCode: fiber.StatusBadRequest}
	}

	// Check Restrictions
	if admin.IsDeviceRestricted {
		isAllowed = false
		type Iot struct {
			BrowserName string `json:"browserName"`
			BrowserId   string `json:"browserId"`
		}
		in := []byte(admin.AllowedDevices.String())
		var iot []Iot
		_ = json.Unmarshal(in, &iot)
		for _, v := range iot {
			if v.BrowserId == payload.BrowserId {
				isAllowed = true
			}
		}
	}

	if admin.IsLocationRestricted {
		type Iot struct {
			Address string  `json:"address"`
			Lat     float64 `json:"lat"`
			Lng     float64 `json:"lng"`
			Type    string  `json:"type"`
		}
		in := []byte(admin.AllowedLocations.String())
		var iot Iot
		_ = json.Unmarshal(in, &iot)
	}

	token := helpers.GenerateToken(strconv.Itoa(int(user.ID)), string(user.Role), user.Name)
	refreshToken := helpers.GenerateRefreshToken(strconv.Itoa(int(user.ID)))
	database.DBSql.Model(&models.Users{}).Where("id=?", user.ID).Updates(models.Users{Token: &token, RefreshToken: &refreshToken})

	if isAllowed {
		return dto.GenericResponse{
			Success: true,
			Data: fiber.Map{
				"user": fiber.Map{
					"name": user.Name,
					"id":   user.ID,
				},
				"admin": fiber.Map{
					"name": admin.Name,
					"id":   admin.ID,
				},
				"token":        token,
				"refreshToken": refreshToken,
			},
			Message:    "Logged-in successfully!",
			StatusCode: fiber.StatusOK,
		}
	} else {
		return dto.GenericResponse{
			Success:    false,
			Data:       nil,
			Message:    "Login restricted",
			StatusCode: fiber.StatusOK,
		}
	}
}
