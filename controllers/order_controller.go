package controllers

import (
	"assignment2/database"
	"assignment2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order = models.Order{
		ID:            c.GetUint("Order_id"),
		Customer_name: c.PostForm("Customer_name"),
		Items: []models.Item{
			{
				Item_code:   c.PostForm("Item_code"),
				Description: c.PostForm("Description"),
				Quantity:    c.GetInt("Quantity"),
			},
		},
		Order_date: c.GetTime("Order_date"),
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.GetDB().Create(&order)
	c.JSON(http.StatusCreated, gin.H{"data": order})
}

func GetOrders(c *gin.Context) {
	var orders []models.Order
	database.GetDB().Preload("Items").Find(&orders)
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func UpdateOrder(c *gin.Context) {
	var order models.Order
	id := c.Param("id")
	if err := database.GetDB().Where("ID = ?", id).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var updatedOrder models.Order
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.GetDB().Model(&updatedOrder).Updates(updatedOrder)
	c.JSON(http.StatusOK, gin.H{"data": order})
}

func DeleteOrder(c *gin.Context) {
	var order models.Order
	id := c.Param("id")
	if err := database.GetDB().Where("ID = ?", id).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	if err := database.GetDB().Select("Items").Delete(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not deleted!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully!"})
	}
}
