package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ClerkID string `bson:"clerkID" json:"clerkID"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName string `bson:"lastName" json:"lastName"`
	Bio *string `bson:"bio" json:"bio"`
	Picture *string `bson:"picture" json:"picture"` 
	Email string `bson:"email" json:"email"`
	EncryptedPassword string `bson:"password" json:"-"`
	Location *string `bson:"location" json:"location"`
	PortfolioWebsite *string `bson:"portfolioWebsite" json:"portfolioWebsite"`
	IsAdmin bool `bson:"isAdmin" json:"isAdmin"`
	Reputation *int `bson:"reputation" json:"reputation"`
	Saved []Question `bson:"saved" json:"saved"`
	JoinedAt time.Time `bson:"joinedAt" json:"joinedAt"`
}

type CreateUserParam struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserParam struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}