package routes

import (
	"coffee-culture.uk/internal/config"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// No version routes
	NoVersionRoutes(router)

	// Version controlled routes
	versionControlled := router.Group("/" + config.AppConfig().App.ApiVersion)
	DefaultRoutes(versionControlled)

}