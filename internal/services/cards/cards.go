package cards

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoyaltyCard struct {
	ID                   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	ShopId               primitive.ObjectID `json:"shopId,omitempty" bson:"shopID,omitempty" validate:"required"`
	CustomerID           primitive.ObjectID `json:"customerId,omitempty" bson:"customerID,omitempty" validate:"required"`
	StampsCollected      int                `json:"stampsCollected,omitempty" bson:"stampsCollected,omitempty"`
	StampsGoal           int                `json:"stampsGoal,omitempty" bson:"stampsGoal,omitempty" validate:"required"`
	TotalStampsCollected int                `json:"totalStampsCollected,omitempty" bson:"totalStampsCollected,omitempty"`
	RewardsCollected     *int               `json:"rewardsCollected,omitempty" bson:"rewardsCollected,omitempty"`
	CreatedAt            primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt            primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Added                bool               `json:"added,omitempty" bson:"added,omitempty"`
}
