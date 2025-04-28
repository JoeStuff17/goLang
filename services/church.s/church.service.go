package church_s

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"main.go/database"
	"main.go/helpers"
	dto "main.go/interface_model"
	"main.go/models"
	sql "main.go/models"
)

func CreateChurch(payload sql.Churches, localUser dto.ReqUser) dto.GenericResponse {
	newUUID := uuid.New().String()
	payload.Uuid = newUUID

	createdBy := dto.CreatedBy{Id: localUser.ID, Name: localUser.Name, Role: localUser.Role}
	createdByJSON, err := json.Marshal(createdBy)
	if err != nil {
		return dto.GenericResponse{Success: false, Data: err.Error(), Message: "Failed to serialize CreatedBy", StatusCode: fiber.StatusUnprocessableEntity}
	}
	payload.CreatedBy = datatypes.JSON(createdByJSON)

	dbWithRetry := helpers.NewDBWithRetry(database.DBSql)
	err = dbWithRetry.CreateWithDynamicGenerator(&payload, func() error {
		newCode, codeErr := GenerateNewChurchCode()
		if codeErr != nil {
			return codeErr
		}
		payload.ChurchCode = newCode
		return nil
	})

	if err != nil {
		return dto.GenericResponse{Success: false, Data: err.Error(), Message: "Failed to create church", StatusCode: fiber.StatusOK}
	}
	return dto.GenericResponse{Success: true, Message: "Church created successfully", Data: payload, StatusCode: fiber.StatusOK}
}

func GenerateNewChurchCode() (string, error) {
	var lastChurchCode string
	err := database.DBSql.Model(&models.Churches{}).
		Select("church_code").
		Where("SUBSTRING(church_code, 3, 2) = ?", time.Now().Format("06")).
		Order("church_code DESC").
		Limit(1).
		Pluck("church_code", &lastChurchCode).Error
	if err != nil {
		return "", err
	}

	codePrefix := "CH"
	yearSuffix := time.Now().Format("06")

	if lastChurchCode == "" {
		return fmt.Sprintf("%s%s001", codePrefix, yearSuffix), nil
	}

	lastNumericPart := lastChurchCode[len(codePrefix)+2:]
	lastNumber, err := strconv.Atoi(lastNumericPart)
	if err != nil {
		return "", err
	}

	newNumber := lastNumber + 1
	newChurchCode := fmt.Sprintf("%s%s%03d", codePrefix, yearSuffix, newNumber)
	return newChurchCode, nil
}
