package dto

import "main.go/enums"

type GenericResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"statusCode"`
}

type ReqUser struct {
	ID   uint           `json:"id"`
	Name string         `json:"name"`
	Role enums.UserRole `json:"role"`
}
