package sql

import (
	"time"

	_ "gorm.io/gorm"
)

const TableNameProductGroup = "product_group"

type ProductGroup struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	GroupName string    `gorm:"column:groupName; type:varchar(255);not null" validate:"required,min=2,max=30" json:"groupName"`
	GroupCode string    `gorm:"column:groupCode; type:varchar(255);not null" validate:"required,min=2,max=30" json:"groupCode"`
	IsActive  *bool     `gorm:"column:isActive; type:bool; NOT NULL DEFAULT TRUE" json:"isActive"`
	CreatedAt time.Time `gorm:"column:createdAt;autoCreateTime;type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;autoUpdateTime;type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (*ProductGroup) TableName() string {
	return TableNameProductGroup
}
