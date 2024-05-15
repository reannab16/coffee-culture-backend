package shop

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ MONGO OBJECTS ------------------------------------------
// ---------------------------------------------------------------------------------------------------
type Shop struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	Password         string             `json:"-" bson:"password,omitempty" validate:"required"`
	Email            string             `json:"email" bson:"email" validate:"required,email"`
	ShopName         *string            `json:"shopName,omitempty" bson:"shopName,omitempty"`
	AdminPassword    *string            `json:"adminPassword,omitempty" bson:"adminPassword,omitempty"`
	Locations        *[]Location        `json:"location,omitempty" bson:"location,omitempty"`
	About            *string            `json:"about,omitempty" bson:"about,omitempty"`
	LoyaltyDeals     *LoyaltyDeals      `json:"loyaltyDeals,omitempty" bson:"loyaltyDeals,omitempty"`
	BackgroundColour *string            `json:"backgroundColour,omitempty" bson:"backgroundColour,omitempty"`
	Images           *Images            `json:"images,omitempty" bson:"images,omitempty"`
}

type Location struct {
	Postcode *string `json:"postcode,omitempty" bson:"postcode,omitempty"`
	Location *string `json:"location,omitempty" bson:"location,omitempty"`
}

type LoyaltyDeals struct {
	Reward     string `json:"reward,omitempty" bson:"reward,omitempty"`
	StampsGoal int    `json:"stampsGoal,omitempty" bson:"stampsGoal,omitempty"`
}

type Images struct {
	Logo    string   `json:"logo,omitempty" bson:"logo,omitempty"`
	MainImg string   `json:"mainImg,omitempty" bson:"mainImg,omitempty"`
	Imgs    []string `json:"imgs,omitempty" bson:"imgs,omitempty"`
}
