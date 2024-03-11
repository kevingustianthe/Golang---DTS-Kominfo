package handlers

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/database"
	"go-rest-api/models"
	"net/http"
	"time"
)

func CreateOrder(c *gin.Context) {
	var reqBody struct {
		CustomerName string        `json:"customer_name"`
		Items        []models.Item `json:"items"`
	}

	err := c.BindJSON(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newOrder := models.Order{
		CustomerName: reqBody.CustomerName,
		OrderedAt:    time.Now(),
		Items:        reqBody.Items,
	}

	database.DB.Create(&newOrder)
	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order": newOrder})
}

func GetOrder(c *gin.Context) {
	var orders []models.Order
	database.DB.Preload("Items").Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func UpdateOrder(c *gin.Context) {
	var reqBody struct {
		CustomerName string        `json:"customer_name"`
		Items        []models.Item `json:"items"`
	}

	err := c.BindJSON(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	orderID := c.Param("id")
	var order models.Order
	var items []models.Item

	err = database.DB.Preload("Items").First(&order, orderID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	order.CustomerName = reqBody.CustomerName
	order.Items = reqBody.Items
	items = reqBody.Items

	database.DB.Save(&order)
	database.DB.Save(&items)
	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully", "order": order})
}

func DeleteOrder(c *gin.Context) {
	orderID := c.Param("id")
	var order models.Order

	err := database.DB.First(&order, orderID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	database.DB.Delete(&order)
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
