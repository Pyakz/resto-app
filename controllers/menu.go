package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}
func RemoveMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}
func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}
func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
