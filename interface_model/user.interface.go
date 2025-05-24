package dto

import "gorm.io/datatypes"

type SendOtpPayload struct {
	MobileNumber string         `json:"mobile_number"`
	Role         string         `json:"role"`
	Location     datatypes.JSON `json:"location"`
}

type VerifyOtpPayload struct {
	MobileNumber string         `json:"mobile_number"`
	Role         string         `json:"role"`
	Location     datatypes.JSON `json:"location"`
	Otp          int            `json:"otp"`
	BrowserId    string         `json:"browser_id"`
}

type FetchUserWithDetails struct {
	UserId     int     `json:"user_id"`
	Role       string  `json:"role"`
	RelatedId  int     `json:"related_id"`
	AdminName  *string `json:"admin_name,omitempty"`
	MemberName *string `json:"member_name,omitempty"`
	ChurchId   *int    `json:"church_id,omitempty"`
	ChurchName *string `json:"church_name,omitempty"`
}
