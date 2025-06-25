package dto

type ChurchFamiliesFetchPayload struct {
	ChurchId int `json:"church_id"`
}

type ChurchFamilyFetchPayload struct {
	ChurchId int `json:"church_id"`
	FamilyId int `json:"family_id"`
}
