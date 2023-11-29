package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/fullstack/dev-overflow/db"
	"github.com/fullstack/dev-overflow/db/fixtures"
	"github.com/fullstack/dev-overflow/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	var (
		mongoDBEndpoint = os.Getenv("MONGO_DB_URL")
		mongoDBName = os.Getenv("MONGO_DB_NAME")
		clerkApiKey = os.Getenv("CLERK_SECRET_KEY")
	)

	clerkClient, err := clerk.NewClient(clerkApiKey)
	if err != nil {
		log.Fatal(err)
	}

	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoDBEndpoint))
	if err != nil {
		log.Fatal(err)
	}

	if err := mongoClient.Database(mongoDBName).Drop(context.Background()); err != nil {
		log.Fatal(err)
	}

	questionStore := db.NewMongoQuestionStore(mongoClient)
	store := &db.Store{
		Question: questionStore,
	}

	clerkUser, err := clerkClient.Users().Read("user_2XUKLyYOZGc5jlzwdFMjyXlHrOw")
	if err != nil {
		log.Fatal(err)
	}

	question := fixtures.AddQuestion(store, "Bagaimana cara mengatur GOROOT", "ini adalah contoh deskripsi", clerkUser.ID, []types.Tag{}, time.Now())
	fmt.Println(question.ID)

}