package models

import (
	"company-service/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID                uuid.UUID         `gorm:"type:uuid;primary_key;" json:"id"`
	Name              string            `gorm:"size:15;unique;not null" json:"name"`
	Description       string            `gorm:"size:3000" json:"description,omitempty"`
	AmountOfEmployees int               `gorm:"not null" json:"amount_of_employees"`
	Registered        bool              `gorm:"not null" json:"registered"`
	Type              utils.CompanyType `sql:"type:ENUM('Corporations', 'NonProfit', 'Cooperative', 'Sole Proprietorship');not null" gorm:"column:company_type" json:"type"`
}

func (company *Company) BeforeCreate(tx *gorm.DB) (err error) {
	company.ID = uuid.New()
	return
}
