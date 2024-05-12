package routes

import (
	"github.com/gin-gonic/gin"
	"coffee-culture.uk/internal/middleware"
	"coffee-culture.uk/internal/services/customer"
)

// GROUP: /user
func CustomerRoutes(group *gin.RouterGroup) {
	// Setup the repository & handler
	customerRepo := customer.NewCustomerRepository()
	customerHandler := customer.NewCustomerHandler(customerRepo)

	customerGroup := group.Group("/customer")

	// --- PUBLIC ROUTES ---
	customerGroup.POST("/create", func(c *gin.Context) {
		customerHandler.CreateCustomer(c)
	})

	// --- PROTECTED ROUTES ---
	customerGroup.Use(middleware.JWTAuthMiddleware())
	{
		customerGroup.GET("/current", func(c *gin.Context) {
			customerHandler.GetCurrentCustomer(c)
		})
	}
}