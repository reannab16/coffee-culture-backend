package customer

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func ConvertCreateCustomerRequestToUser(req CreateCustomerRequest) (Customer, error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return Customer{}, err
	}

	return Customer{
		ID:          primitive.NewObjectID(),
		Username:    req.Username,
		Password:    string(hashed),
		Email:       req.Email,
	}, nil
}