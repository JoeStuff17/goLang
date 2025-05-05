package church_user_s

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
	"main.go/database"
	"main.go/enums"
	dto "main.go/interface_model"
	"main.go/models"
)

func CreateChurchUser(payload *models.ChurchUser, localUser dto.ReqUser) dto.GenericResponse {
	// var churchDetails models.Churches
	// err1 := database.DBSql.Model(&models.Churches{}).Where("id = ?", payload.ChurchId).First(&churchDetails)
	// if err1 != nil {
	// 	return dto.GenericResponse{Success: false, Message: "Church not found", Data: err1, StatusCode: fiber.StatusNotFound}
	// }

	createdBy := dto.CreatedBy{Id: localUser.ID, Name: localUser.Name, Role: localUser.Role}
	createdByJSON, err := json.Marshal(createdBy)
	if err != nil {
		return dto.GenericResponse{Success: false, Message: "Failed to serialize CreatedBy", Data: err.Error(), StatusCode: fiber.StatusUnprocessableEntity}
	}
	payload.CreatedBy = datatypes.JSON(createdByJSON)

	// if churchDetails.AdminRoles.
	err = database.DBSql.Model(&models.ChurchUser{}).Create(&payload).Error
	if err != nil {
		return dto.GenericResponse{Success: false, Message: "Failed to create church", Data: err.Error(), StatusCode: fiber.StatusInternalServerError}
	}
	var userRole enums.UserRole
	if enums.UserRole(payload.Designation) == enums.UserRole(enums.ChurchUserRoleMember) {
		userRole = enums.RoleChurchUser
	} else {
		userRole = enums.RoleChurchAdmin
	}
	userPayload := &models.Users{
		Name:         payload.Name,
		MobileNumber: payload.MobileNumber,
		Role:         userRole,
		RelatedId:    uint16(payload.ID),
	}
	err = database.DBSql.Model(&models.Users{}).Create(&userPayload).Error
	if err != nil {
		return dto.GenericResponse{Success: false, Message: "Failed to create user", Data: err.Error()}
	}

	return dto.GenericResponse{
		Success:    true,
		Message:    "User created successfully",
		Data:       &payload,
		StatusCode: fiber.StatusOK,
	}
}

func GetAllChurchUsers(church_id int) dto.ResWithCount {
	var members []models.Churches
	err := database.DBSql.Where("church_id = ?", church_id).Find(&members).Error
	if err != nil {
		return dto.ResWithCount{Success: false, Message: err.Error(), Data: []map[string]interface{}{}, Count: 0,
			StatusCode: fiber.StatusNoContent,
		}
	}
	message := "No members found"
	if len(members) > 0 {
		message = "Members fetched successfully"
	}
	return dto.ResWithCount{Success: true, Message: message, Data: &members, Count: len(members), StatusCode: fiber.StatusOK}
}

func FetchChurchUserById(church_id int, user_id int) dto.GenericResponse {
	var user *models.ChurchUser
	result := database.DBSql.Where("church_id = ? AND id = ? AND is_active = ?", church_id, user_id, true).Find(&user)
	if result.RowsAffected == 0 {
		return dto.GenericResponse{Success: true, Data: nil, Message: "No records found", StatusCode: fiber.StatusOK}
	}
	return dto.GenericResponse{Success: true, Message: "User retrieved successfully", Data: user, StatusCode: fiber.StatusOK}
}
