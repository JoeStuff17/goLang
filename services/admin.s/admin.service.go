package admin_s

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"main.go/database"
	"main.go/enums"
	dto "main.go/interface_model"
	"main.go/models"
)

func CreateAdmin(payload models.Admins, localUser dto.ReqUser) dto.GenericResponse {
	var adminCount, userCount int64
	database.DBSql.Model(&models.Admins{}).Where("mobile_number = ?", payload.MobileNumber).Count(&adminCount)
	database.DBSql.Model(&models.Users{}).Where("mobile_number = ? AND role = ?", payload.MobileNumber, enums.RoleAdmin).Count(&userCount)
	if userCount == 0 && adminCount == 0 {
		newUUID := uuid.New().String()
		EmployeeCode, _ := GenerateEmployeeCode()
		createdBy := dto.CreatedBy{Id: localUser.ID, Name: localUser.Name, Role: localUser.Role}
		createdByJSON, err1 := json.Marshal(createdBy)
		if err1 != nil {
			return dto.GenericResponse{
				Success: false,
				Data:    err1.Error(),
			}
		}
		payload.CreatedBy = datatypes.JSON(createdByJSON)
		adminPayload := &models.Admins{
			Name:         payload.Name,
			MobileNumber: payload.MobileNumber,
			Designation:  payload.Designation,
			CreatedBy:    payload.CreatedBy,
			EmployeeCode: EmployeeCode,
			UUID:         newUUID,
		}
		err := database.DBSql.Model(&models.Admins{}).Create(&adminPayload).Error
		if err != nil {
			return dto.GenericResponse{Success: false, Message: err.Error(), Data: nil}
		}
		// create user
		userPayload := &models.Users{
			Name:         payload.Name,
			MobileNumber: payload.MobileNumber,
			Role:         enums.RoleAdmin,
			RelatedId:    uint16(adminPayload.ID),
		}
		_ = database.DBSql.Model(&models.Users{}).Create(&userPayload).Error
		if err != nil {
			return dto.GenericResponse{
				Success: false,
				Data:    err.Error(),
			}
		}
		return dto.GenericResponse{
			Success:    true,
			Data:       nil,
			Message:    "Admin Created successfully",
			StatusCode: fiber.StatusOK,
		}
	} else {
		return dto.GenericResponse{
			Success:    false,
			Data:       nil,
			Message:    "Admin already exists",
			StatusCode: fiber.StatusConflict,
		}
	}
}

func GenerateEmployeeCode() (string, error) {
	var lastEmployeeCode string
	err := database.DBSql.Model(&models.Admins{}).Select("employee_code").Order("employee_code DESC").Limit(1).Pluck("employee_code", &lastEmployeeCode).Error
	if err != nil {
		return "", err
	}
	if lastEmployeeCode == "" {
		return "GM0001", nil
	}

	codePrefix := "GM"
	lastNumericPart := lastEmployeeCode[len(codePrefix):]
	lastNumber, err := strconv.Atoi(lastNumericPart)
	if err != nil {
		return "", err
	}
	newNumber := lastNumber + 1
	newEmployeeCode := fmt.Sprintf("%s%04d", codePrefix, newNumber)
	return newEmployeeCode, nil
}

func FetchAllAdmins() dto.GenericResponse {
	var admins []map[string]interface{}
	database.DBSql.Raw("SELECT id, name, email, mobile_number, role, designation, employee_code, status, allowed_locations, allowed_devices from admins").Scan(&admins)
	if len(admins) == 0 {
		return dto.GenericResponse{
			Success:    true,
			Data:       []map[string]interface{}{},
			Message:    "No admins found",
			StatusCode: fiber.StatusOK,
		}
	}
	return dto.GenericResponse{
		Success:    true,
		Message:    "Admins retrieved successfully",
		Data:       admins,
		StatusCode: fiber.StatusOK,
	}
}

func FetchAdminById(admin_id int) dto.GenericResponse {
	var admin *models.Admins
	result := database.DBSql.Where("id = ?", admin_id).Find(&admin)
	if result.RowsAffected == 0 {
		return dto.GenericResponse{Success: true, Data: nil, Message: "No records found", StatusCode: fiber.StatusOK}
	}
	return dto.GenericResponse{Success: true, Message: "Admin retrieved successfully", Data: admin, StatusCode: fiber.StatusOK}
}

func UpdateAdminMenu(payload *dto.UpdateAdminMenuPayload) dto.GenericResponse {
	var adminCount int64
	database.DBSql.Model(&models.Admins{}).Where("id = ?", payload.ID).Count(&adminCount)
	if adminCount == 0 {
		return dto.GenericResponse{Success: false, Data: nil, Message: "Admin is not active", StatusCode: fiber.StatusBadRequest}
	}
	database.DBSql.Model(&models.Admins{}).Where("id =?", payload.ID).Updates(models.Admins{Menu: payload.Data.Menu})
	return dto.GenericResponse{Success: true, Data: fiber.Map{}, Message: "Fetched successfully", StatusCode: fiber.StatusOK}
}
