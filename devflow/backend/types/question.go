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
	Tags []Tag `bson:"tags" json:"tags"`
	Views int `bson:"views" json:"views"`
	Upvotes int `bson:"upvotes" json:"upvotes"`
	Downvotes int `bson:"downvotes" json:"downvotes"`
	Answers []Answer `bson:"answers" json:"answers"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

type AskQuestionParams struct {
	Title string `json:"title"`
	Description string `json:"description"`
	UserID primitive.ObjectID `json:"useID"`
	Tags []string `json:"tags"`
	CreatedAt time.Time `json:"createdAt"`
}