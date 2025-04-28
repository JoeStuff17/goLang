package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const TableNameChurches = "churches"

type Churches struct {
	Name           string                         `gorm:"column:name; type:varchar(255);not null" json:"name"`
	ChurchCode     string                         `gorm:"column:church_code; unique; type:varchar(20); not null" json:"church_code"`
	Uuid           string                         `gorm:"column:uuid; type:varchar(255);not null" json:"uuid"`
	LogoUrl        *string                        `gorm:"column:logo_url; type:varchar(255);" json:"logo_url"`
	CoverImageUrl  *string                        `gorm:"column:cover_image_url; type:varchar(255);not null" json:"cover_image_url"`
	Address        string                         `gorm:"column:address; type:varchar(255);not null" json:"address"`
	DistrictId     int16                          `gorm:"column:district_id; type:uint;not null" json:"district_id"`
	ZipCode        string                         `gorm:"column:zipcode; type:varchar(255); not null" json:"zipcode"`
	Email          *string                        `gorm:"column:email; type:varchar(255)" json:"email"`
	MobileNumber   *string                        `gorm:"column:mobile_number; type:varchar(255)" json:"mobile_number"`
	Website        *string                        `gorm:"column:website; type:varchar(255)" json:"website"`
	InauguratedAt  datatypes.Date                 `gorm:"column:inaugurated_at; type:DATETIME; not null" json:"inaugurated_at"`
	DenominationId int16                          `gorm:"column:denomination_id; type:uint" json:"denomination_id"`
	ClientIds      *datatypes.JSONArrayExpression `gorm:"column:client_ids; type:text" json:"client_ids"`
	CreatedBy      datatypes.JSON                 `gorm:"column:created_by; type:text" json:"created_by"`
	IsActive       bool                           `gorm:"column:is_active; default:true" json:"is_active"`
	gorm.Model
}

func (*Churches) TableName() string {
	return TableNameChurches
}

func (c *Churches) BeforeSave(tx *gorm.DB) error {
	c.Name = strings.TrimSpace(c.Name)
	if c.MobileNumber != nil {
		trimmed := strings.TrimSpace(*c.MobileNumber)
		c.MobileNumber = &trimmed
	}
	if c.Email != nil {
		trimmed := strings.TrimSpace(*c.Email)
		c.Email = &trimmed
	}
	return nil
}

func (c *Churches) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ChurchCode == "" {
		newCode, err := GenerateChurchCode(tx)
		if err != nil {
			return err
		}
		c.ChurchCode = newCode
	}
	return nil
}

// func GenerateChurchCode(tx *gorm.DB) (string, error) {
// 	var maxChurchCode string
// 	yearSuffix := time.Now().Format("06")
// 	codePrefix := "CH"

// 	err := tx.
// 		Model(&Churches{}).
// 		Select("MAX(church_code)").
// 		Where("SUBSTRING(church_code, 3, 2) = ?", yearSuffix).
// 		Clauses(clause.Locking{Strength: "UPDATE"}). // lock to prevent concurrency issue(race condition)
// 		Scan(&maxChurchCode).Error
// 	if err != nil {
// 		return "", err
// 	}

// 	if maxChurchCode == "" {
// 		return fmt.Sprintf("%s%s001", codePrefix, yearSuffix), nil
// 	}

// 	lastNumericPart := maxChurchCode[len(codePrefix)+2:]
// 	lastNumber, err := strconv.Atoi(lastNumericPart)
// 	if err != nil {
// 		return "", err
// 	}

// 	newNumber := lastNumber + 1
// 	newChurchCode := fmt.Sprintf("%s%s%03d", codePrefix, yearSuffix, newNumber)
// 	return newChurchCode, nil
// }

func GenerateChurchCode(tx *gorm.DB) (string, error) {
	var maxChurchCode *string
	yearSuffix := time.Now().Format("06")
	codePrefix := "CH"

	row := tx.
		Model(&Churches{}).
		Select("MAX(church_code)").
		Where("SUBSTRING(church_code, 3, 2) = ?", yearSuffix).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Row()

	err := row.Scan(&maxChurchCode)
	if err != nil {
		return "error ------> ", err
	}

	if maxChurchCode == nil || *maxChurchCode == "" {
		return fmt.Sprintf("%s%s001", codePrefix, yearSuffix), nil
	}

	if len(*maxChurchCode) < 5 {
		return "", fmt.Errorf("invalid existing church code format: %s", *maxChurchCode)
	}

	lastNumericPart := (*maxChurchCode)[len(codePrefix)+2:] // skip prefix + year
	lastNumber, err := strconv.Atoi(lastNumericPart)
	if err != nil {
		return "", err
	}

	newNumber := lastNumber + 1
	newChurchCode := fmt.Sprintf("%s%s%03d", codePrefix, yearSuffix, newNumber)
	return newChurchCode, nil
}
