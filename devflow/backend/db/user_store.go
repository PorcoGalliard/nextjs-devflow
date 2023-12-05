package db

import (
	"context"
	"os"

	"github.com/fullstack/dev-overflow/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const USERCOLL = "users"

type MongoUserStore struct {
	client *mongo.Client
	collection *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	var mongoEnvDBName = os.Getenv("MONGO_DB_NAME")
	return &MongoUserStore{
		client: client,
		collection: client.Database(mongoEnvDBName).Collection(USERCOLL),
	}
}

type UserStore interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
	GetUserByID(context.Context, string) (*types.User, error) 
	UpdateUser(context.Context, string, *types.UpdateUserParam) (*types.User, error)
}

func (s *MongoUserStore) CreateUser(c context.Context, user *types.User) (*types.User, error) {
	res, err := s.collection.InsertOne(c, user)
	if err != nil {
		return nil , err
	}
	
	user.ID = res.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	if err := s.collection.FindOne(ctx, bson.M{"clerkID": id}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, clerkID string, update *types.UpdateUserParam) (*types.User ,error) {

	filter := bson.M{"clerkID": clerkID}
	updateData := bson.M{"$set": update.UpdateData}

	result := s.collection.FindOneAndUpdate(ctx, filter, updateData)

	var updatedUser types.User
	if err := result.Decode(&updatedUser); err != nil {
		return nil, err
	}

	return &updatedUser, nil
}