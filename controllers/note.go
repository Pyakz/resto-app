package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func GetNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}
func RemoveNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}
func UpdateNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}
func CreateNote() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
