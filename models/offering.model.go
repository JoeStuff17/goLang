package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"main.go/enums"
)

const TableNameOfferings = "offerings"

type Offerings struct {
	ChurchId     int                  `gorm:"column:church_id; not null" json:"church_id"`
	MemberId     int                  `gorm:"column:member_id; not null" json:"member_id"`
	Amount       float64              `gorm:"column:amount; type:decimal(13,2); NOT NULL;DEFAULT:0" json:"amount"`
	OfferingMode enums.PaymentMethod  `gorm:"column:offering_mode;type:enum('cash','online','bank_transfer','cheque'); DEFAULT: 'cash'" json:"offering_mode"`
	Purpose      enums.PaymentPurpose `gorm:"column:purpose;type:enum('tithe','offering','donation','building_fund','thanksgiving','vbs','sunday_school','youth_fellowship', 'women_fellowship', 'men_fellowship', 'auction', 'evangelistic_fund', 'imm'); DEFAULT: 'tithe'" json:"purpose"`
	ForMonth     enums.Months         `gorm:"column:for_month;type:enum('1','2','3','4','5','6','7','8','9','10','11','12'); not null" json:"for_month"`
	ForYear      int16                `gorm:"column:for_year;not null" json:"for_year"`
	CreatedBy    datatypes.JSON       `gorm:"column:created_by; type:text" json:"created_by"`
	gorm.Model
}

func (*Offerings) TableName() string {
	return TableNameOfferings
}
