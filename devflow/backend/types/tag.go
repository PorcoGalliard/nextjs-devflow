package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Description string `bson:"description" json:"description"`
	Name string `bson:"name" json:"name"`
	Questions []Question `bson:"questions" json:"questions"`
	Followers []User `bson:"followers" json:"followers"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

type CreateTagParams struct {
	Name string `json:"name"`
}

type UpdateTagQuestionAndFollowers struct {
	Questions []Question `json:"questions"`
	Followers []User `json:"followers"`
}