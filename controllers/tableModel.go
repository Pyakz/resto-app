package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}
func RemoveTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}
func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}
func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
