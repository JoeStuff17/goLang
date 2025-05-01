package church_s

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
	"main.go/database"
	"main.go/helpers"
	dto "main.go/interface_model"
	"main.go/models"
)

func CreateChurch(payload *models.Churches, localUser dto.ReqUser) dto.GenericResponse {
	createdBy := dto.CreatedBy{
		Id:   localUser.ID,
		Name: localUser.Name,
		Role: localUser.Role,
	}
	createdByJSON, err := json.Marshal(createdBy)
	if err != nil {
		return dto.GenericResponse{
			Success:    false,
			Message:    "Failed to serialize CreatedBy",
			Data:       err.Error(),
			StatusCode: fiber.StatusUnprocessableEntity,
		}
	}
	payload.CreatedBy = datatypes.JSON(createdByJSON)
	dbRetry := helpers.NewDBWithRetry(database.DBSql)
	err = dbRetry.CreateWithDynamicGenerator(payload, func() error {
		return payload.BeforeCreate(database.DBSql)
	})
	if err != nil {
		return dto.GenericResponse{
			Success:    false,
			Message:    "Failed to create church",
			Data:       err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		}
	}

	return dto.GenericResponse{
		Success:    true,
		Message:    "Church created successfully",
		Data:       &payload,
		StatusCode: fiber.StatusOK,
	}
}

func GetAllChurches() dto.ResWithCount {
	var churches []models.Churches
	err := database.DBSql.Model(&models.Churches{}).Find(&churches).Error
	if err != nil {
		return dto.ResWithCount{
			Success:    false,
			Message:    err.Error(),
			Data:       []map[string]interface{}{},
			Count:      0,
			StatusCode: fiber.StatusNoContent,
		}
	}
	message := "No churches found"
	if len(churches) > 0 {
		message = "Churches fetched successfully"
	}
	return dto.ResWithCount{
		Success:    true,
		Message:    message,
		Data:       &churches,
		Count:      len(churches),
		StatusCode: fiber.StatusOK,
	}
}
