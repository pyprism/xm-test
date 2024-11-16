package main

import (
	"company-service/models"
	"company-service/routes"
)

func main() {
	models.ConnectDatabase()
	routes.Init()
}
