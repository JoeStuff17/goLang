package dto

type OfferingsFetchPayload struct {
	ChurchId  int  `query:"church_id"`
	MemberId  *int `query:"member_id"`
	FromMonth *int `query:"from_month"`
	FromYear  *int `query:"from_year"`
	ToMonth   *int `query:"to_month"`
	ToYear    *int `query:"to_year"`
}

// member access
type FetchOfferingsByMemberPayload struct {
	FromMonth *int `query:"from_month"`
	FromYear  *int `query:"from_year"`
	ToMonth   *int `query:"to_month"`
	ToYear    *int `query:"to_year"`
}
