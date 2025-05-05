package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const TableNameChurchFamilies = "church_families"

type ChurchFamily struct {
	ChurchId       string         `gorm:"column:church_id; type:varchar(20); not null" json:"church_id"`
	FamilyName     string         `gorm:"column:family_name; type:varchar(255);not null" json:"family_name"`
	MemberIds      datatypes.JSON `gorm:"column:member_ids; type:text" json:"member_ids"`
	TotalMembers   int            `gorm:"column:total_members; type:int; not null" json:"total_members"`
	FamilyImageUrl *string        `gorm:"column:family_image_url; type:varchar(255);" json:"family_image_url"`
	CreatedBy      datatypes.JSON `gorm:"column:created_by; type:text" json:"created_by"`
	IsActive       bool           `gorm:"column:is_active; default:true" json:"is_active"`
	gorm.Model
}

func (*ChurchFamily) TableName() string {
	return TableNameChurchFamilies
}
