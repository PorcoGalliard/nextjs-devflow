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
	coll *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	var mongoEnvDBName = os.Getenv("MONGO_DB_NAME")
	return &MongoUserStore{
		client: client,
		coll: client.Database(mongoEnvDBName).Collection(USERCOLL),
	}
}

type UserStore interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
	GetUserByID(context.Context, string) (*types.User, error) 
	UpdateUser(context.Context, string, *types.UpdateUserParam) (*types.User, error)
	DeleteUser(context.Context, string) error
}

func (s *Store) CreateUser(c context.Context, user *types.User) (*types.User, error) {
	res, err := s.User.coll.InsertOne(c, user)
	if err != nil {
		return nil , err
	}
	
	user.ID = res.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (s *Store) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	if err := s.User.coll.FindOne(ctx, bson.M{"clerkID": id}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Store) UpdateUser(ctx context.Context, clerkID string, update *types.UpdateUserParam) (*types.User ,error) {

	filter := bson.M{"clerkID": clerkID}
	updateData := bson.M{"$set": update.UpdateData}

	result := s.User.coll.FindOneAndUpdate(ctx, filter, updateData)

	var updatedUser types.User
	if err := result.Decode(&updatedUser); err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (s *Store) DeleteUser(ctx context.Context, clerkID string) error {
	user, err := s.UserStore.GetUserByID(ctx, clerkID)
	if err != nil {
		return err
	}

	_, err = s.Question.coll.DeleteMany(ctx, bson.M{"userID": user.ID})
	if err != nil {
		return err
	}

	_, err = s.Tag.collection.UpdateMany(ctx, bson.M{"followers": user.ID}, bson.M{"$pull": bson.M{"followers": user.ID}})
	if err != nil {
		return err
	}

	_, err = s.User.coll.DeleteOne(ctx, bson.M{"userID": user.ID})
	if err != nil {
		return err
	}
	
	return nil
}