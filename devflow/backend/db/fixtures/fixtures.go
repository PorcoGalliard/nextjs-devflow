package fixtures

import (
	"context"
	"log"
	"time"

	"github.com/fullstack/dev-overflow/db"
	"github.com/fullstack/dev-overflow/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddQuestion(store *db.Store, title string, desc string, userID primitive.ObjectID, tags []types.Tag, createdAt time.Time) (*types.Question) {

	question := &types.Question{
		Title: title,
		Description: desc,
		UserID: userID,
		Tags: tags,
		CreatedAt: createdAt,
	}

	insertedQuestion, err := store.Question.AskQuestion(context.Background(), question)
	if err != nil {
		log.Fatal(err)
	}

	return insertedQuestion
}

func AddUser(store *db.Store, firstName string, lastName string, clerkID string, email string, encryptedPassword string) (*types.User) {
	user, err := types.NewUserFromParams(types.CreateUserParam{
		FirstName: firstName,
		LastName: lastName,
		ClerkID: clerkID,
		Email: email,
		Password: encryptedPassword,
	})

	if err != nil {
		log.Fatal(err)
	}

	insertedUser, err := store.User.CreateUser(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	return insertedUser
}