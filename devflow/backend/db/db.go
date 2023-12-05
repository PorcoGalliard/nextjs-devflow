package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const MongoDBName = "MONGO_DB_NAME"

type Store struct {
	Question *MongoQuestionStore
	User *MongoUserStore
	Tag *MongoTagStore
	QuestionStore QuestionStore
	UserStore UserStore
	TagStore TagStore
}

func NewStore(client *mongo.Client) *Store {
	return &Store{
		Question: NewMongoQuestionStore(client),
		User: NewMongoUserStore(client),
		Tag: NewMongoTagStore(client),
	}
}