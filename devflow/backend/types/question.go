package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	UserID string `bson:"userID,omitempty" json:"userID,omitempty"`
	Tags []string `bson:"tags" json:"tags"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

type AskQuestionParams struct {
	Title string `json:"title"`
	Description string `json:"description"`
	UserID primitive.ObjectID `json:"useID"`
	Tags []string `json:"tags"`
	CreatedAt time.Time `json:"createdAt"`
}