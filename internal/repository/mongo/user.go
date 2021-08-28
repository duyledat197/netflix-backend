package mongo

import (
	"context"
	"log"

	"github.com/duyledat197/netfix-backend/common/errors"
	"github.com/duyledat197/netfix-backend/internal/models"
	"github.com/duyledat197/netfix-backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	coll *mongo.Collection
}

func (r *userRepository) Create(user *models.User) error {
	ctx := context.Background()
	err := user.HashPassword()
	if err != nil {
		return err
	}
	_, err = r.coll.InsertOne(ctx, user)
	return err
}

func (r *userRepository) UpdateInfo(userID string, user *models.User) error {
	ctx := context.Background()
	_, err := r.coll.UpdateOne(ctx, bson.M{"_id": userID}, user)
	return err
}

func (r *userRepository) UpdatePassword(userID, password string) error {
	ctx := context.Background()
	query := bson.M{
		"$set": bson.M{"password": password},
	}
	_, err := r.coll.UpdateOne(ctx, bson.M{"_id": userID}, query)
	return err
}

func (r *userRepository) FindByID(userID string) (*models.User, error) {
	ctx := context.Background()
	result := &models.User{}
	err := r.coll.FindOne(ctx, bson.M{"_id": userID}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	ctx := context.Background()
	result := &models.User{}
	err := r.coll.FindOne(ctx, bson.M{"email": email}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, errors.ErrEmailNotFound
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) FindAll(offset int64, limit int64) ([]*models.User, error) {

	ctx := context.Background()
	results := []*models.User{}

	cur, err := r.coll.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result *models.User
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

// NewUserRepository ...
func NewUserRepository(coll *mongo.Collection) repository.UserRepository {
	return &userRepository{
		coll: coll,
	}
}
