package fixtures

import (
	"context"
	"log"
	"time"

	"github.com/fullstack/dev-overflow/db"
	"github.com/fullstack/dev-overflow/types"
)

func AddQuestion(store *db.Store, title string, desc string, userID string, tags []string, createdAt time.Time) (*types.Question) {
	question := &types.Question{
		Title: title,
		Description: desc,
		UserID: userID,
		Tags: []string(tags),
		CreatedAt: createdAt,
	}

	insertedQuestion, err := store.Question.AskQuestion(context.Background(), question)
	if err != nil {
		log.Fatal(err)
	}

	return insertedQuestion
}