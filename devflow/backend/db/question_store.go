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
	TagStore
	UserStore
}

func NewMongoQuestionStore(client *mongo.Client, tagStore TagStore, userStore UserStore) *MongoQuestionStore {
	var mongoenvdbname = os.Getenv("MONGO_DB_NAME")
	return &MongoQuestionStore{
		client: client,
		coll: client.Database(mongoenvdbname).Collection(QUESTIONCOLL),
		TagStore: tagStore,
		UserStore: userStore,
	}
}

type QuestionStore interface {
	Dropper
	GetQuestionByID(context.Context, string) (*types.Question, error)
	GetQuestionsByUserID(context.Context, string) ([]*types.Question, error)
	GetQuestions(context.Context) ([]*types.Question, error)
	AskQuestion(context.Context, *types.Question) (*types.Question, error)
	DeleteQuestionByID(context.Context, string) error
	DeleteManyQuestionsByUserID(context.Context, primitive.ObjectID) error
}

func (s *MongoQuestionStore) Drop(ctx context.Context) error {
	fmt.Println("****DELETING DATABASE****")
	return s.coll.Drop(ctx)
}

func (s *MongoQuestionStore) GetQuestionByID(ctx context.Context, id string) (*types.Question, error) {

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var question types.Question
	if err := s.coll.FindOne(ctx, bson.M{"_id":oid}).Decode(&question); err != nil {
		return nil , err
	}

	return &question, nil
}

func (s *MongoQuestionStore) GetQuestionsByUserID(ctx context.Context, id string) ([]*types.Question, error) {
	var questions []*types.Question

	user, err := s.UserStore.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	oid := user.ID

	cursor, err := s.coll.Find(ctx, bson.M{"userID": oid})
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

	return questions, nil

}

func (s *MongoQuestionStore) GetQuestions(ctx context.Context) ([]*types.Question, error) {
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

	cursor, err := s.coll.Aggregate(ctx, pipeline)
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

func (s *MongoQuestionStore) AskQuestion(ctx context.Context, question *types.Question) (*types.Question, error) {
	res, err := s.coll.InsertOne(ctx, question)
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

func (s *MongoQuestionStore) DeleteQuestionByID(ctx context.Context, id string) error {

	oid,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if err := s.coll.FindOneAndDelete(ctx, bson.M{"_id": oid}).Err(); err != nil {
		return err
	}
	return nil
}

func (s *MongoQuestionStore) DeleteManyQuestionsByUserID(ctx context.Context, id primitive.ObjectID) error {
	_, err := s.coll.DeleteMany(ctx, bson.M{"userID": id})
	if err != nil {
		return err
	}
	return nil
}