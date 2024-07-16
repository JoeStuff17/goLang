package sql

import (
	"time"

	_ "gorm.io/gorm"
)

const TableNameProductType = "product_type"

type ProductType struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	GroupName string    `gorm:"column:groupName; type:varchar(255);not null" validate:"required,min=2,max=30" json:"groupName"`
	GroupCode string    `gorm:"column:groupCode; type:varchar(255);not null" validate:"required,min=2,max=30" json:"groupCode"`
	IsActive  bool      `gorm:"column:isActive; default:true" json:"isActive"`
	CreatedAt time.Time `gorm:"column:createdAt;autoCreateTime;type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;autoUpdateTime;type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (*ProductType) TableName() string {
	return TableNameProductType
}
