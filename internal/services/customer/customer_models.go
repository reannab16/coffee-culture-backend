package customer

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ MONGO OBJECTS ------------------------------------------
// ---------------------------------------------------------------------------------------------------
type Customer struct {
	ID          primitive.ObjectID    `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	Username    string                `json:"username,omitempty" bson:"username,omitempty" validate:"required"`
	Password    string                `json:"-" bson:"password,omitempty" validate:"required"`
	Email       string                `json:"email" bson:"email" validate:"required,email"`
	FirstName   *string               `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName    *string               `json:"lastName,omitempty" bson:"lastName,omitempty"`
	LoyaltyCard []*primitive.ObjectID `json:"university,omitempty" bson:"university,omitempty"`
	Location    *string               `json:"location,omitempty" bson:"location,omitempty"`
	CreatedAt   primitive.DateTime    `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}


