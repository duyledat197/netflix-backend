package mongo

import (
	"context"
	"log"

	"github.com/duyledat197/netfix-backend/internal/models"
	"github.com/duyledat197/netfix-backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type categoryRepository struct {
	coll *mongo.Collection
}

// NewCategoryRepository ...
func NewCategoryRepository(coll *mongo.Collection) repository.CategoryRepository {
	return &categoryRepository{
		coll: coll,
	}
}

func (r *categoryRepository) Create(category *models.Category) error {
	ctx := context.Background()
	_, err := r.coll.InsertOne(ctx, category)
	return err
}
func (r *categoryRepository) Delete(id primitive.ObjectID) error {
	ctx := context.Background()

	if _, err := r.coll.DeleteOne(ctx, &models.Category{ID: id}); err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) FindAll() ([]*models.Category, error) {
	ctx := context.Background()
	results := []*models.Category{}

	cur, err := r.coll.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result *models.Category
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
		// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *categoryRepository) FindByID(id primitive.ObjectID) (*models.Category, error) {
	ctx := context.Background()
	result := &models.Category{}

	if err := r.coll.
		FindOne(ctx, &models.Category{
			ID: id,
		}).
		Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
