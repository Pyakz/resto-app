package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}
func RemoveOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}
func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
