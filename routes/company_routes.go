package routes

import (
	"company-service/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	companyRoutes := router.Group("/companies")
	{
		companyRoutes.POST("", controllers.CreateCompany)
		companyRoutes.GET("", controllers.ListCompany)
		companyRoutes.PATCH("/:id", controllers.PatchCompanyById)
		companyRoutes.DELETE("/:id", controllers.DeleteCompanyById)
		companyRoutes.GET("/:id", controllers.GetCompanyById)
	}

	return router
}
