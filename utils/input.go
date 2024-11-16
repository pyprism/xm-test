package utils

type PostCompanyRequest struct {
	Name              string      `json:"name" binding:"required"`
	Description       string      `json:"description"`
	AmountOfEmployees int         `json:"amount_of_employees" binding:"required"`
	Registered        bool        `json:"registered"`
	Type              CompanyType `json:"type" binding:"required"`
}
