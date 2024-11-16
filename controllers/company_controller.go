package controllers

import (
	"company-service/services"
	"company-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func CreateCompany(c *gin.Context) {
	var company utils.PostCompanyRequest
	bindErr := c.BindJSON(&company)

	// basic input validations
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "POST body parse error", "error": bindErr.Error()})
		return
	}
	if company.Type != utils.Corporations && company.Type != utils.NonProfit && company.Type != utils.Sole {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid company type", "type": company.Type})
		return
	}

	// save data
	if err := services.CreateCompanyFromRequest(&company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "create company error", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Company has been created", "data": company})
}

func ListCompany(context *gin.Context) {
	companies, err := services.GetAllCompanies()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "list companies error", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": companies})
}

func GetCompanyById(context *gin.Context) {
	companyId := context.Param("id")
	companyIdUUID, err := uuid.Parse(companyId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "invalid uuid", "error": err.Error()})
		return
	}
	company, err := services.GetCompanyById(companyIdUUID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "get company by id error", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": company})
}

func PatchCompanyById(context *gin.Context) {
	companyId := context.Param("id")
	companyIdUUID, err := uuid.Parse(companyId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "invalid uuid", "error": err.Error()})
		return
	}
	company, err := services.GetCompanyById(companyIdUUID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "get company by id error", "error": err.Error()})
		return
	}

	var updateData map[string]interface{}
	if err := context.ShouldBindJSON(&updateData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := services.PatchCompany(company, updateData); err != nil {
		context.JSON(http.StatusNotModified, gin.H{"message": "Failed to update company", "error": err.Error()})
		return
	}
	log.Println("Company updated id: ", companyId)
	context.JSON(http.StatusOK, gin.H{"message": "Company has been updated", "data": company})
}

func DeleteCompanyById(context *gin.Context) {
	companyId := context.Param("id")
	companyIdUUID, err := uuid.Parse(companyId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "invalid uuid", "error": err.Error()})
	}
	err = services.DeleteCompanyById(companyIdUUID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Delete company error", "error": err.Error()})
		return
	}
	log.Println("Company deleted id: ", companyId)
	context.JSON(http.StatusOK, gin.H{"message": "Company has been deleted"})

}
