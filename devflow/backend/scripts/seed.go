package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fullstack/dev-overflow/db"
	"github.com/fullstack/dev-overflow/db/fixtures"

	// "github.com/fullstack/dev-overflow/types"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	tagStore := db.NewMongoTagStore(mongoClient)
	store := &db.Store{
		Question: questionStore,
		User: userStore,
		Tag: tagStore,
	}


	user := fixtures.AddUser(store, "Higuruma", "Hiromi", "12345678", "higuruma@gmail.com", "higurumahiromu123")
	fmt.Println("User berhasil ditambahkan, berikut adalah ID user =>", user.ID)

	user2 := fixtures.AddUser(store, "Yuuta", "Okkotsu", "87654321", "okkotsu@gmail.com", "yuutaokkotsu123")
	fmt.Println("User berhasil ditambahkan, berikut adalah ID user =>", user2.ID)

	user3 := fixtures.AddUser(store, "Geto", "Suguru", "55554444", "getosuguru@gmail.com", "getosugu123")
	fmt.Println("User berhasil ditambahkan, berikut adalah ID user =>", user3.ID)

	tag := fixtures.AddTag(store, "Golang")

	taggg := make([]primitive.ObjectID, 0)
	taggg = append(taggg, tag.ID)
	question := fixtures.AddQuestion(store, "Bagaimana cara mengatur GOROOT", "ini adalah contoh deskripsi", user.ID, taggg, time.Now().UTC())
	fmt.Println("Pertanyaan berhasil ditambahkan, berikut adalah ID pertanyaan =>", question.ID)

	// questions := []primitive.ObjectID{question.ID, question2.ID}
	// followers := []primitive.ObjectID{user.ID}

	// updatedTag, err := fixtures.UpdateTag(store, tag.ID, &types.UpdateTagQuestionAndFollowers{Questions: questions, Followers: followers})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Tag berhasil diupdate =>", updatedTag.ID)
}