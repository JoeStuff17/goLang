package dto

import (
	"gorm.io/datatypes"
	"main.go/enums"
)

type CreatedBy struct {
	Id   uint           `json:"id"`
	Name string         `json:"name"`
	Role enums.UserRole `json:"role"`
}

type UpdateAdminMenuPayload struct {
	ID   uint          `json:"id"`
	Data AdminMenuData `json:"data"`
}

type AdminMenuData struct {
	Menu datatypes.JSON `json:"menu"`
}
