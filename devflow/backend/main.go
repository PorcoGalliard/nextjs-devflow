package main

import (
	"context"
	"log"
	"os"

	"github.com/fullstack/dev-overflow/api"
	"github.com/fullstack/dev-overflow/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoEndpoint := os.Getenv("MONGO_DB_URL")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoEndpoint))
	if err != nil {
		log.Fatal(err)
	} 

	var (
		questionStore = db.NewMongoQuestionStore(client)

		store = &db.Store{
			Question: questionStore,
		}

		questionHandler = api.NewQuestionHandler(store.Question)
		app = fiber.New()
		apiv1 = app.Group("/api/v1")
	)

	apiv1.Get("/question/:id", questionHandler.HandleGetQuestionByID)
	

	listenAddr := os.Getenv("HTTP_LISTEN_ADDRESS")
	app.Listen(listenAddr)
}