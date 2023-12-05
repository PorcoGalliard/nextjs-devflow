package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fullstack/dev-overflow/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const QUESTIONCOLL = "questions"

type Map map[string]any

type Dropper interface {
	Drop(context.Context) error
}

type MongoQuestionStore struct {
	client *mongo.Client
	coll *mongo.Collection
}

func NewMongoQuestionStore(client *mongo.Client) *MongoQuestionStore {
	var mongoenvdbname = os.Getenv("MONGO_DB_NAME")
	return &MongoQuestionStore{
		client: client,
		coll: client.Database(mongoenvdbname).Collection(QUESTIONCOLL),
	}
}

type QuestionStore interface {
	Dropper
	GetQuestionByID(context.Context, string) (*types.Question, error)
	GetQuestions(context.Context) ([]*types.Question, error)
	AskQuestion(context.Context, *types.Question) (*types.Question, error)
}

func (s *Store) Drop(ctx context.Context) error {
	fmt.Println("****DELETING DATABASE****")
	return s.Question.coll.Drop(ctx)
}

func (s *Store) GetQuestionByID(ctx context.Context, id string) (*types.Question, error) {

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var question types.Question
	if err := s.Question.coll.FindOne(ctx, bson.M{"_id":oid}).Decode(&question); err != nil {
		return nil , err
	}

	return &question, nil
}

func (s *Store) GetQuestions(ctx context.Context) ([]*types.Question, error) {
	var questions []*types.Question

	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from": "users",
				"localField": "userID",
				"foreignField": "_id",
				"as": "user",
			}},
			{"$unwind":"$user"},
			{"$lookup":bson.M{
				"from": "tags",
				"localField": "tags",
				"foreignField": "_id",
				"as": "tagDetails",
			}},
			{"$sort":bson.M{"createdAt":-1}},
	}

	cursor, err := s.Question.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var question types.Question
		if err := cursor.Decode(&question); err != nil {
			return nil, err
		}

		questions = append(questions, &question)
	}

	log.Println(questions)
	return questions, nil
}

func (s *Store) AskQuestion(ctx context.Context, question *types.Question) (*types.Question, error) {
	res, err := s.Question.coll.InsertOne(ctx, question)
	if err != nil {
		return nil, err
	}

	question.ID = res.InsertedID.(primitive.ObjectID)

	for _, tag := range question.Tags {
		if err := s.TagStore.UpdateTag(ctx, Map{"_id": tag}, &types.UpdateTagQuestionAndFollowers{Questions: question.ID, Followers: question.UserID}); err != nil {
			return nil, err
		}
	}

	return question, nil
}