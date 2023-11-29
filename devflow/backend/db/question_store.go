package db

import (
	"context"
	"fmt"
	"os"

	"github.com/fullstack/dev-overflow/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const QuestionColl = "questions"

type Map map[string]any

type Dropper interface {
	Drop(context.Context)
}

type MongoQuestionStore struct {
	client *mongo.Client
	coll *mongo.Collection
}

func NewMongoQuestionStore(client *mongo.Client) *MongoQuestionStore {
	var mongoenvdbname = os.Getenv("MONGO_DB_NAME")
	return &MongoQuestionStore{
		client: client,
		coll: client.Database(mongoenvdbname).Collection(QuestionColl),
	}
}

type QuestionStore interface {
	GetQuestionByID(context.Context, string) (*types.Question, error)
	AskQuestion(context.Context, *types.Question) (*types.Question, error)
}

func (s *MongoQuestionStore) Drop(ctx context.Context) error {
	fmt.Println("****DELETING DATABASE****")
	return s.coll.Drop(ctx)
}

func (s *MongoQuestionStore) GetQuestionByID(ctx context.Context, id string) (*types.Question, error) {
	var question types.Question
	if err := s.coll.FindOne(ctx, bson.M{"_id":id}).Decode(&question); err != nil {
		return nil , err
	}

	return &question, nil
}

func (s *MongoQuestionStore) AskQuestion(ctx context.Context, question *types.Question) (*types.Question, error) {
	res, err := s.coll.InsertOne(ctx, question)
	if err != nil {
		return nil, err
	}
	question.ID = res.InsertedID.(primitive.ObjectID)
	return question, nil
}
