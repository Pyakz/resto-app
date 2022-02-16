package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}
func RemoveOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}
func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}
func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func GetOrderItemsByOrderId() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}
