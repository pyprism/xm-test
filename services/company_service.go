package services

import (
	"company-service/models"
	"company-service/utils"
	"github.com/google/uuid"
)

func CreateCompanyFromRequest(companyData *utils.PostCompanyRequest) error {
	company := models.Company{
		Name:              companyData.Name,
		Description:       companyData.Description,
		AmountOfEmployees: companyData.AmountOfEmployees,
		Registered:        companyData.Registered,
		Type:              companyData.Type,
	}
	if err := CreateCompany(company); err != nil {
		return err
	}
	return nil
}

func CreateCompany(company models.Company) error {
	return models.DB.Create(&company).Error
}

func PatchCompany(company *models.Company, updateData map[string]interface{}) error {
	return models.DB.Model(&company).Updates(updateData).Error
}

func GetCompanyById(id uuid.UUID) (*models.Company, error) {
	var company models.Company
	err := models.DB.Where("id = ?", id).First(&company).Error
	if err != nil {
		return nil, err

	}
	return &company, nil
}

func DeleteCompanyById(id uuid.UUID) error {
	var company models.Company
	err := models.DB.Where("id = ?", id).Delete(&company).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllCompanies() ([]models.Company, error) {
	var companies []models.Company
	err := models.DB.Find(&companies).Error
	if err != nil {
		return nil, err
	}
	return companies, nil
}
