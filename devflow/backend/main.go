package main

import (
	"context"
	"log"
	"os"

	"github.com/fullstack/dev-overflow/api"
	"github.com/fullstack/dev-overflow/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	config = fiber.Config{
		ErrorHandler: api.ErrorHandler,
	}
)

func main() {
	mongoEndpoint := os.Getenv("MONGO_DB_URL")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoEndpoint))
	if err != nil {
		log.Fatal(err)
	} 

	var (
		questionStore = db.NewMongoQuestionStore(client)
		userStore = db.NewMongoUserStore(client)

		store = &db.Store{
			Question: questionStore,
			User: userStore,
		}

		questionHandler = api.NewQuestionHandler(store.Question, store.User)
		userHandler = api.NewUserHandler(store.User)
		app = fiber.New(config)
		auth = app.Group("/api")
		apiv1 = app.Group("/api/v1")
	)

	app.Use(cors.New())
	// Question Handler
	apiv1.Get("/question/:_id", questionHandler.HandleGetQuestionByID)
	apiv1.Post("/ask-question", questionHandler.HandleAskQuestion)

	// User Handler
	auth.Post("/sign-up", userHandler.HandleSignUp)
	apiv1.Get("/user/:clerkID", userHandler.HandleGetUserByID)
	

	listenAddr := os.Getenv("HTTP_LISTEN_ADDRESS")
	app.Listen(listenAddr)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}