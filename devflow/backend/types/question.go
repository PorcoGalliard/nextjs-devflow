package types

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	minTitleLength = 20
	minDescriptionLength = 100
	maxTagsLength = 3
)

type Question struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	UserID primitive.ObjectID `bson:"userID,omitempty" json:"userID,omitempty"`
	Tags []primitive.ObjectID `bson:"tags" json:"tags"`
	Views int `bson:"views" json:"views"`
	Upvotes []primitive.ObjectID `bson:"upvotes" json:"upvotes"`
	Downvotes []primitive.ObjectID `bson:"downvotes" json:"downvotes"`
	Answers []primitive.ObjectID `bson:"answers" json:"answers"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

type AskQuestionParams struct {
	Title string `json:"title"`
	Description string `json:"description"`
	ClerkID string `json:"clerkID"`
	Tags []string `json:"tags"`
}

func (params AskQuestionParams) Validate() map[string]string {
	errors := map[string]string{}

	if len(params.Title) < minTitleLength {
		errors["title"] = fmt.Sprintf("Title should be at least %d characters", minTitleLength)
	}

	if len(params.Description) < minDescriptionLength {
		errors["description"] = fmt.Sprintf("Description must be at least %d characters", minDescriptionLength)
	}

	if len(params.Tags) > maxTagsLength {
		errors["tags"] = fmt.Sprintf("Tag must be at most %d items", maxTagsLength)
	}

	return errors
}