package dto

type UpdatePayload struct {
	ID   uint                    `json:"id"`
	Data *map[string]interface{} `json:"data"`
}
