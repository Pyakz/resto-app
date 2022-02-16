package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Hello %s", "World")
	}
}
func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Food ID: %s", id)
	}
}
func RemoveFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Food ID: %s", id)

	}
}
func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Food ID: %s", id)

	}
}
func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
