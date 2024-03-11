package main

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/database"
	"go-rest-api/routes"
)

func main() {
	database.Connect()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8082")
}
