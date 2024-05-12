package customer

import (
	"github.com/gin-gonic/gin"
	// "golang.org/x/crypto/bcrypt"
	"net/http"
	"coffee-culture.uk/internal/api"
	// "coffee-culture.uk/internal/middleware"
)

type CustomerHandler struct {
	Repo CustomerRepository
}

func NewCustomerHandler(repo CustomerRepository) *CustomerHandler {
	return &CustomerHandler{Repo: repo}
}

// @Summary Create a new customer
// @Description Create a new customer
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body CreateCustomerRequest true "Customer object to be created"
// @Success 201 {object} CustomerResponse "Customer created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /customer/create [post]
func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var reqPayload CreateCustomerRequest
	if err := c.ShouldBindJSON(&reqPayload); err != nil {
		api.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	mongoCustomer, err := ConvertCreateCustomerRequestToUser(reqPayload)
	if err != nil {
		api.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	createdCustomer, err := h.Repo.CreateCustomer(c.Request.Context(), mongoCustomer)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.Success(c, http.StatusCreated, "Created user successfully", createdCustomer)
}

func (h *CustomerHandler) GetCurrentCustomer(c *gin.Context) {
	
}