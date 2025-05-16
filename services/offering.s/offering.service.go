package offerings_s

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
	"main.go/database"
	dto "main.go/interface_model"
	"main.go/models"
)

func CreateOffering(payload *models.Offerings, localUser dto.ReqUser) dto.GenericResponse {
	createdBy := dto.CreatedBy{Id: localUser.ID, Name: localUser.Name, Role: localUser.Role}
	createdByJSON, err := json.Marshal(createdBy)
	if err != nil {
		return dto.GenericResponse{Success: false, Message: "Failed to serialize CreatedBy", Data: err.Error(), StatusCode: fiber.StatusUnprocessableEntity}
	}
	payload.CreatedBy = datatypes.JSON(createdByJSON)

	err = database.DBSql.Model(&models.Offerings{}).Create(&payload).Error
	if err != nil {
		return dto.GenericResponse{Success: false, Message: "Failed to add offering", Data: err.Error(), StatusCode: fiber.StatusInternalServerError}
	}

	return dto.GenericResponse{Success: true, Message: "offerings added successfully", Data: &payload, StatusCode: fiber.StatusOK}
}

func FetchChurchOfferings(church_id int) dto.ResWithCount {
	var offerings []models.Offerings
	println("church_id", church_id)
	err := database.DBSql.Where("church_id = ?", church_id).Find(&offerings).Error
	if err != nil {
		return dto.ResWithCount{Success: false, Message: err.Error(), Data: []map[string]interface{}{}, Count: 0,
			StatusCode: fiber.StatusNoContent,
		}
	}
	message := "Offerings not found"
	if len(offerings) > 0 {
		message = "Offerings fetched successfully"
	}
	return dto.ResWithCount{Success: true, Message: message, Data: &offerings, Count: len(offerings), StatusCode: fiber.StatusOK}
}

func FetchOfferingsByMember(church_id int, member_id int) dto.GenericResponse {
	var user *models.Offerings
	result := database.DBSql.Where("church_id = ? AND id = ?", church_id, member_id).Find(&user)
	if result.RowsAffected == 0 {
		return dto.GenericResponse{Success: true, Data: nil, Message: "Offerings not found", StatusCode: fiber.StatusOK}
	}
	return dto.GenericResponse{Success: true, Message: "Offerings retrieved successfully", Data: user, StatusCode: fiber.StatusOK}
}
