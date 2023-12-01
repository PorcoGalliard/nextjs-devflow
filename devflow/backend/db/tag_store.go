package db

import (
	"context"
	"os"

	"github.com/fullstack/dev-overflow/types"
	"github.com/fullstack/dev-overflow/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const TAGCOLL = "tags"

type MongoTagStore struct {
	client *mongo.Client
	collection *mongo.Collection
}

func NewMongoTagStore(client *mongo.Client) *MongoTagStore {
	var mongoEnvDBName = os.Getenv("MONGO_DB_NAME")
	return &MongoTagStore{
		client: client,
		collection: client.Database(mongoEnvDBName).Collection(TAGCOLL),
	}
}

type TagStore interface {
	CreateTag(context.Context, *types.Tag) (*types.Tag, error)
	GetTagByID(context.Context, string) (*types.Tag, error)
	GetTagByName(context.Context, string) (*types.Tag, error)
}

func (s *MongoTagStore) GetTagByID(ctx context.Context, id string) (*types.Tag, error) {

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var tag types.Tag
	if err := s.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&tag); err != nil {
		return nil, err
	}

	return &tag, nil
}

func (s *MongoTagStore) GetTagByName(ctx context.Context, name string) (*types.Tag, error) {
	var tag types.Tag
	name = utils.FormatTag(name)
	if err := s.collection.FindOne(ctx, bson.M{"name": name}).Decode(&tag); err != nil {
		return nil, err
	}

	return &tag, nil
}

func (s *MongoTagStore) CreateTag(c context.Context, tag *types.Tag) (*types.Tag, error) {
	tag.Name = utils.FormatTag(tag.Name)
	res, err := s.collection.InsertOne(c, tag)
	if err != nil {
		return nil , err
	}
	
	tag.ID = res.InsertedID.(primitive.ObjectID)

	return tag, nil
}