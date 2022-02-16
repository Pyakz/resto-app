package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}
func RemoveInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}
func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
