package productGroup

import (
	"main.go/database"
	"main.go/enums"
	sql "main.go/models"
)

func CreateProductGroup(payload *sql.ProductGroup) enums.GenericResponse {

	err := database.DBSql.Db.Create(payload).Error
	if err != nil {
		return enums.GenericResponse{
			Success: false,
			Data:    err,
		}
	}
	return enums.GenericResponse{
		Success: true,
		Data:    payload.ID,
	}
}
