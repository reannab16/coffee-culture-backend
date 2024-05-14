package customer

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"coffee-culture.uk/internal/api"
	"net/http"
	"coffee-culture.uk/internal/middleware"
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

// @Summary Login customer
// @Description Login customer
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body CustomerLoginRequest true "Customer login credentials"
// @Success 200 {object} CustomerLoginResponse "Customer logged in successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 401 {object} map[string]interface{} "Invalid credentials"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /customer/login [post]
func (h *CustomerHandler) LoginUser(c *gin.Context) {
	var reqPayload CustomerLoginRequest
	if err := c.ShouldBindJSON(&reqPayload); err != nil {
		api.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	customer, err := h.Repo.GetCustomerByEmail(c.Request.Context(), reqPayload.Email)
	if err != nil {
		api.Error(c, http.StatusNotFound, "Customer not found")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(reqPayload.Password)); err != nil {
		api.Error(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := middleware.GenerateJWTToken(customer.Email, customer.Username)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, "An error occurred while processing your request")
		return
	}

	customerRes := CustomerResponse{
		ID:          customer.ID.Hex(),
		Username:    customer.Username,
		Email:       customer.Email,
		
	}

	api.Success(c, http.StatusOK, "Customer logged in successfully", CustomerLoginResponse{
		Token: token,
		Customer:  customerRes,
	})
}