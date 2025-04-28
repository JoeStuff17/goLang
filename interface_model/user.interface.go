package dto

import "gorm.io/datatypes"

type AdminSendOtpPayload struct {
	MobileNumber string         `json:"mobile_number"`
	Location     datatypes.JSON `json:"location"`
}

type AdminVerifyOtpPayload struct {
	MobileNumber string         `json:"mobile_number"`
	Location     datatypes.JSON `json:"location"`
	Otp          int            `json:"otp"`
	BrowserId    string         `json:"browser_id"`
}
