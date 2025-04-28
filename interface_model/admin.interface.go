package dto

import "main.go/enums"

type CreatedBy struct {
	Id   uint           `json:"id"`
	Name string         `json:"name"`
	Role enums.UserRole `json:"role"`
}
