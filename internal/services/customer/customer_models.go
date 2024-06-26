package customer

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// ---------------------------------------------------------------------------------------------------
// ----------------------------------------- CREATE OBJECTS ----------------------------------------
// ---------------------------------------------------------------------------------------------------

type CreateCustomerRequest struct {
	// ID          string  `json:"id,omitempty"`
	Username    string  `json:"username" binding:"required,alphanum,min=3,max=30"`
	Email       string  `json:"email" binding:"required,email"`
	Password    string  `json:"password" binding:"required,min=6"`

}

// ---------------------------------------------------------------------------------------------------
// ----------------------------------------- RESPONSE OBJECTS ----------------------------------------
// ---------------------------------------------------------------------------------------------------
type CustomerResponse struct {
	ID        string `json:"id,omitempty"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type CustomerLoginResponse struct {
	Token string       `json:"token"`
	Customer  CustomerResponse `json:"customer"`
}

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ MONGO OBJECTS ------------------------------------------
// ---------------------------------------------------------------------------------------------------
type Customer struct {
	ID           primitive.ObjectID    `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	Username     string                `json:"username,omitempty" bson:"username,omitempty" validate:"required"`
	Password     string                `json:"-" bson:"password,omitempty" validate:"required"`
	Email        string                `json:"email" bson:"email" validate:"required,email"`
	FirstName    *string               `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName     *string               `json:"lastName,omitempty" bson:"lastName,omitempty"`
	LoyaltyCard  []*primitive.ObjectID `json:"university,omitempty" bson:"university,omitempty"`
	Location     *string               `json:"location,omitempty" bson:"location,omitempty"`
	StampHistory *[]Stamp              `json:"stampHistory,omitempty" bson:"stampHistory,omitempty"`
	CreatedAt    primitive.DateTime    `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

type Stamp struct {
	Date      primitive.DateTime `json:"date,omitempty" bson:"date,omitempty"`
	ShopId    string             `json:"shopId,omitempty" bson:"shopId,omitempty"`
	StampType string             `json:"stampType,omitempty" bson:"stampType,omitempty"`
}
