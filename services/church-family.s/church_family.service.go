package church_family_s

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
	"main.go/database"
	dto "main.go/interface_model"
	"main.go/models"
)

func CreateChurchFamily(payload *models.ChurchFamily, localUser dto.ReqUser) dto.GenericResponse {
	var churchFamily models.ChurchFamily
	database.DBSql.Model(&models.ChurchFamily{}).Where("JSON_CONTAINS(member_ids, ?)", fmt.Sprintf("[%d]", localUser.ID)).First(&churchFamily)
	if churchFamily.ID != 0 {
		// family already exists - update the family
		membersList := churchFamily.MemberIds
		membersList = append(membersList, payload.MemberIds...)
		payload.MemberIds = membersList
		database.DBSql.Model(&models.ChurchFamily{}).Where("id =?", churchFamily.ID).Updates(models.ChurchFamily{MemberIds: payload.MemberIds})
		return dto.GenericResponse{Success: true, Message: "Family updated successfully", Data: &payload, StatusCode: fiber.StatusOK}
	}

	randomDigits := GenerateRandomString(5)
	payload.FamilyName = "family" + randomDigits
	// payload.MemberIds, _ = json.Marshal([]uint{localUser.ID, payload.MemberIds...})
	createdBy := dto.CreatedBy{Id: localUser.ID, Name: localUser.Name, Role: localUser.Role}
	createdByJSON, err := json.Marshal(createdBy)
	if err != nil {
		return dto.GenericResponse{Success: false, Message: "Failed to serialize CreatedBy", Data: err.Error(), StatusCode: fiber.StatusUnprocessableEntity}
	}
	payload.CreatedBy = datatypes.JSON(createdByJSON)
	err = database.DBSql.Model(&models.ChurchFamily{}).Create(&payload).Error
	if err != nil {
		return dto.GenericResponse{Success: false, Message: "Failed to create church family", Data: err.Error(), StatusCode: fiber.StatusInternalServerError}
	}

	return dto.GenericResponse{
		Success:    true,
		Message:    "User created successfully",
		Data:       &payload,
		StatusCode: fiber.StatusOK,
	}
}

func GenerateRandomString(length int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func GetAllChurchFamilies(church_id int) dto.ResWithCount {
	var families []models.ChurchFamily
	err := database.DBSql.Where("church_id = ? AND is_active = ?", church_id, true).Find(&families).Error
	if err != nil {
		return dto.ResWithCount{Success: false, Message: err.Error(), Data: []map[string]interface{}{}, Count: 0,
			StatusCode: fiber.StatusNoContent,
		}
	}
	message := "Family details not found"
	if len(families) > 0 {
		message = "families fetched successfully"
	}
	return dto.ResWithCount{Success: true, Message: message, Data: &families, Count: len(families), StatusCode: fiber.StatusOK}
}

func FetchChurchFamilyById(church_id int, family_id int) dto.GenericResponse {
	var family *models.ChurchFamily
	result := database.DBSql.Where("church_id = ? AND id = ? AND is_active = ?", church_id, family_id, true).Find(&family)
	if result.RowsAffected == 0 {
		return dto.GenericResponse{Success: true, Data: nil, Message: "Family details not found", StatusCode: fiber.StatusOK}
	}
	return dto.GenericResponse{Success: true, Message: "family retrieved successfully", Data: family, StatusCode: fiber.StatusOK}
}
