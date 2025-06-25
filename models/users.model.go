package models

import (
	"database/sql"
	"strings"

	"main.go/enums"

	"gorm.io/gorm"
)

const TableNameUsers = "users"

type Users struct {
	Name          string         `gorm:"column:name; type:varchar(100); NOT NULL" json:"name"`
	Role          enums.UserRole `gorm:"column:role;type:enum('super_admin','admin','church','church_admin','church_user'); DEFAULT: 'admin'" json:"role"`
	MobileNumber  string         `gorm:"column:mobile_number; type:varchar(10);" json:"mobile_number"`
	Email         *string        `gorm:"column:email; type:varchar(150);" json:"email"`
	OTP           int            `gorm:"column:otp; type:varchar(6);" json:"otp"`
	Token         *string        `gorm:"column:token; type:varchar(255);" json:"token"`
	RefreshToken  *string        `gorm:"column:refresh_token; type:varchar(255);" json:"refresh_token"`
	RelatedId     uint16         `gorm:"column:related_id; type:uint" json:"related_id"` // if role admin then related id is admin id
	LastLoginTime sql.NullTime   `gorm:"column:lastLoginTime;type:TIMESTAMP; DEFAULT:null" json:"lastLoginTime"`
	LastOTPSent   sql.NullTime   `gorm:"column:lastOTPSent;type:TIMESTAMP; DEFAULT:null" json:"lastOTPSent"`
	Device        *string        `gorm:"column:device;type:varchar(255)" json:"device"`
	gorm.Model
}

func (*Users) TableName() string {
	return TableNameUsers
}

func (c *Users) BeforeSave(tx *gorm.DB) error {
	c.Name = strings.TrimSpace(c.Name)
	c.MobileNumber = strings.TrimSpace(c.MobileNumber)
	if c.Email != nil {
		trimmed := strings.TrimSpace(*c.Email)
		c.Email = &trimmed
	}
	return nil
}
