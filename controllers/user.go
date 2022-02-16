package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Params ID: %s", id)
	}
}

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func SignUpUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
