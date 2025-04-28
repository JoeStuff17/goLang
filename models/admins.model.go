package models

import (
	"strings"

	"main.go/enums"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const TableNameAdmins = "admins"

type Admins struct {
	Name                 string                 `gorm:"column:name; type:varchar(100); NOT NULL" json:"name"`
	UUID                 string                 `gorm:"column:uuid;type:text" json:"uuid"`
	Role                 enums.UserRole         `gorm:"column:role;type:enum('super_admin','admin','client','client_admin','client_user'); DEFAULT: 'admin'" json:"role"`
	MobileNumber         string                 `gorm:"column:mobile_number; type:varchar(10);" json:"mobile_number"`
	Email                *string                `gorm:"column:email; type:varchar(150);" json:"email"`
	EmployeeCode         string                 `gorm:"column:employee_code; unique; type:varchar(20); not null" validate:"required,min=6,max=6" json:"employee_code"`
	Designation          enums.AdminDesignation `gorm:"column:designation; type:enum('manager','admin','lead','employee'); DEFAULT: 'admin'" json:"designation"`
	Menu                 datatypes.JSON         `gorm:"column:menu;type:text" json:"menu"`
	IsLocationRestricted bool                   `gorm:"column:is_location_restricted; type:TINYINT; NOT NULL; DEFAULT:FALSE" json:"is_location_restricted"`
	IsDeviceRestricted   bool                   `gorm:"column:is_device_restricted; type:TINYINT; NOT NULL; DEFAULT:FALSE" json:"is_device_restricted"`
	AllowedLocations     *datatypes.JSON        `gorm:"column:allowed_locations;type:text" json:"allowed_locations"`
	AllowedDevices       *datatypes.JSON        `gorm:"column:allowed_devices;type:text" json:"allowed_devices"`
	WebFcmToken          *string                `gorm:"column:web_fcm_token; type:text" json:"web_fcm_token"`
	CreatedBy            datatypes.JSON         `gorm:"column:created_by; type:json" json:"created_by"`
	Status               enums.GenericStatus    `gorm:"column:status;type:enum('0','1','2','3');DEFAULT: 1" json:"status"`
	gorm.Model
}

func (*Admins) TableName() string {
	return TableNameAdmins
}

func (c *Admins) BeforeSave(tx *gorm.DB) error {
	c.Name = strings.TrimSpace(c.Name)
	c.MobileNumber = strings.TrimSpace(c.MobileNumber)
	if c.Email != nil {
		trimmed := strings.TrimSpace(*c.Email)
		c.Email = &trimmed
	}
	return nil
}
