package models

import (
	"strings"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"main.go/enums"
)

const TableNameChurchUsers = "church_users"

type ChurchUser struct {
	ChurchId       string                  `gorm:"column:church_id; type:varchar(20); not null" json:"church_id"`
	Name           string                  `gorm:"column:name; type:varchar(255);not null" json:"name"`
	MobileNumber   string                  `gorm:"column:mobile_number; type:varchar(10);" json:"mobile_number"`
	Email          *string                 `gorm:"column:email; type:varchar(150);" json:"email"`
	Designation    enums.ChurchDesignation `gorm:"column:designation; type:enum('pastor','assistant_pastor','priest','admin','treasurer','secretary','elder','deacon','youth_leader','choir_leader','sunday_school_teacher','member'); not null; DEFAULT: 'member'" json:"designation"`
	Gender         enums.Gender            `gorm:"column:gender; type:enum('male','female', 'other'); not null; DEFAULT: 'male'" json:"gender"`
	Dob            *time.Time              `gorm:"column:dob; type:DATETIME; not null" json:"dob"`
	MaritalStatus  enums.MaritalStatus     `gorm:"column:marital_status; type:enum('single','married','divorced','widowed'); not null; DEFAULT: 'single'" json:"marital_status"`
	MarriedAt      *time.Time              `gorm:"column:married_at; type:DATETIME;" json:"married_at"`
	Profession     enums.Professions       `gorm:"column:profession; type:enum('none','student','unemployed','government_employee','private_employee','business_owner','retired','self_employed','homemaker'); not null; DEFAULT: 'none'" json:"profession"`
	LanguagesKnown datatypes.JSON          `gorm:"column:languages_known; type:text" json:"languages_known"`
	FamilyMembers  *datatypes.JSON         `gorm:"column:family_members; type:text" json:"family_members"` // {member_id: 123, relationship: 'spouse'}
	FamilyId       *int                    `gorm:"column:family_id; type:varchar(20);" json:"family_id"`
	Address        *string                 `gorm:"column:address; type:varchar(255);not null" json:"address"`
	DistrictId     int16                   `gorm:"column:district_id; type:uint;not null" json:"district_id"`
	ZipCode        string                  `gorm:"column:zipcode; type:varchar(255); not null" json:"zipcode"`
	IsPoc          bool                    `gorm:"column:is_poc; type:tinyint(1); default:false" json:"is_poc"`
	IsAdmin        bool                    `gorm:"column:is_admin; type:tinyint(1); default:false" json:"is_admin"`
	IsBaptised     bool                    `gorm:"column:is_baptised; type:tinyint(1); default:false" json:"is_baptised"`
	BaptisedAt     *time.Time              `gorm:"column:baptised_at; type:DATETIME;" json:"baptised_at"`
	JoinedAt       *time.Time              `gorm:"column:joined_at; type:DATETIME;" json:"joined_at"`
	TitheGiver     bool                    `gorm:"column:tithe_giver; type:tinyint(1); default:false" json:"tithe_giver"`
	TitheAmount    float64                 `gorm:"column:tithe_amount; type:float; not null; default:0" json:"tithe_amount"`
	CreatedBy      datatypes.JSON          `gorm:"column:created_by; type:text" json:"created_by"`
	IsActive       bool                    `gorm:"column:is_active; default:true" json:"is_active"`
	gorm.Model
}

func (*ChurchUser) TableName() string {
	return TableNameChurchUsers
}

func (c *ChurchUser) BeforeSave(tx *gorm.DB) error {
	c.Name = strings.TrimSpace(c.Name)
	c.MobileNumber = strings.TrimSpace(c.MobileNumber)
	if c.Email != nil {
		trimmed := strings.TrimSpace(*c.Email)
		c.Email = &trimmed
	}
	if c.Email != nil {
		trimmed := strings.TrimSpace(*c.Email)
		c.Email = &trimmed
	}
	return nil
}
