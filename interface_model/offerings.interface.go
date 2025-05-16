package dto

type OfferingsFetchPayload struct {
	ChurchId int `json:"church_id"`
}

type FetchOfferingsByMemberPayload struct {
	ChurchId int `json:"church_id"`
	MemberId int `json:"member_id"`
}
