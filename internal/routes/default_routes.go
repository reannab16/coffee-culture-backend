package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"coffee-culture.uk/internal/api"
	"coffee-culture.uk/internal/config"
)

// For /v0/
func DefaultRoutes(group *gin.RouterGroup) {
	group.GET("/", func(c *gin.Context) {
		api.Success(c, http.StatusOK, "Coffee Culture", gin.H{
			"message": "Welcome to the Coffee Culture REST API v0",
			"version": config.AppConfig().App.AppVersion,
			"author":  "Reanna Bakshani",
		})
	})
}

// For /
func NoVersionRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		api.Success(c, http.StatusOK, "Coffee Culture", gin.H{
			"message": "Welcome to the Coffee Culture REST API v0",
			"version": config.AppConfig().App.AppVersion,
			"author":  "Reanna Bakshani",
			"api":     "v0",
		})
	})
}