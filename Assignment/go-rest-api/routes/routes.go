package routes

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/handlers"
)

func SetupRoutes(router *gin.Engine) {
	orderRoutes := router.Group("/orders")
	{
		orderRoutes.GET("/", handlers.GetOrder)
		orderRoutes.POST("/", handlers.CreateOrder)
		orderRoutes.PUT("/:id", handlers.UpdateOrder)
		orderRoutes.DELETE("/:id", handlers.DeleteOrder)
	}
}
