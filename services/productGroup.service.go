package productGroup

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"main.go/database"
	"main.go/enums"
	"main.go/interface_model"
	sql "main.go/models"
)

var ProductTypeColl = database.OpenConnection(database.Client, "product_type")

func CreateProductGroup(payload *sql.ProductGroup) enums.GenericResponse {

	err := database.DBSql.Model(&sql.ProductGroup{}).Create(payload).Error
	if err != nil {
		return enums.GenericResponse{
			Success: false,
			Data:    err,
		}
	}
	return enums.GenericResponse{
		Success: true,
		Data:    payload,
	}
}

func FetchAllGroups() enums.GenericResponse {
	var groups []sql.ProductGroup
	// err := database.DBSql.Db.Limit(2).Offset(0).Find(&groups).Error
	err := database.DBSql.Model(&sql.ProductGroup{}).Find(&groups).Error
	if err != nil {
		return enums.GenericResponse{
			Success: false,
			Data:    err,
		}
	}
	return enums.GenericResponse{
		Success: true,
		Data:    groups,
	}
}

func UpdateById(payload *interface_model.UpdatePayload) enums.GenericResponse {
	var groups sql.ProductGroup
	err := database.DBSql.Model(&sql.ProductGroup{}).Where("id = ?", payload.ID).Updates(payload.Data).Error
	if err != nil {
		return enums.GenericResponse{
			Success: false,
			Data:    err,
		}
	}
	return enums.GenericResponse{
		Success: true,
		Data:    groups,
	}
}

func CreateInMongo(payload fiber.Map) enums.GenericResponse {

	_, err := ProductTypeColl.InsertOne(context.TODO(), payload)
	if err != nil {
		return enums.GenericResponse{Success: false, Data: err, Message: "unable to update log", StatusCode: fiber.StatusUnprocessableEntity}
	}
	return enums.GenericResponse{Success: true, Data: payload, Message: "updated successfully", StatusCode: fiber.StatusOK}
}
