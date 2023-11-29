package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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
	)

	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoDBEndpoint))
	if err != nil {
		log.Fatal(err)
	}

	if err := mongoClient.Database(mongoDBName).Drop(context.Background()); err != nil {
		log.Fatal(err)
	}

	questionStore := db.NewMongoQuestionStore(mongoClient)
	userStore := db.NewMongoUserStore(mongoClient)
	store := &db.Store{
		Question: questionStore,
		User: userStore,
	}

	user := fixtures.AddUser(store, "Higuruma", "Hiromu", "higuruma@gmail.com", "higurumahiromu123")
	fmt.Println("User berhasil ditambahkan, berikut adalah ID user =>", user.ID)

	question := fixtures.AddQuestion(store, "Bagaimana cara mengatur GOROOT", "ini adalah contoh deskripsi", user.ID, []types.Tag{}, time.Now())
	fmt.Println("Pertanyaan berhasil ditambahkan, berikut adalah ID pertanyaan =>", question.ID)
}