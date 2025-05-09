package dto

type ChurchUsersFetchPayload struct {
	ChurchId int `json:"church_id"`
}

type ChurchUserFetchPayload struct {
	ChurchId int `json:"church_id"`
	UserId   int `json:"user_id"`
}
