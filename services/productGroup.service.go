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
		Data:    payload,
	}
}

func FetchAllGroups() enums.GenericResponse {
	var groups []sql.ProductGroup
	// err := database.DBSql.Db.Limit(2).Offset(0).Find(&groups).Error
	err := database.DBSql.Db.Find(&groups).Error
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

// func UpdateById(payload *sql.ProductGroup) enums.GenericResponse {
// 	var groups []sql.ProductGroup
// 	err := database.DBSql.Db.Update(&payload.IsActive, 1).Error
// 	if err != nil {
// 		return enums.GenericResponse{
// 			Success: false,
// 			Data:    err,
// 		}
// 	}
// 	return enums.GenericResponse{
// 		Success: true,
// 		Data:    groups,
// 	}
// }
