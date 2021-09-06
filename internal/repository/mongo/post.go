package mongo

import (
	"context"
	"log"

	"github.com/duyledat197/netfix-backend/internal/models"
	"github.com/duyledat197/netfix-backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type postRepository struct {
	coll *mongo.Collection
}

// NewPostRepository ...
func NewPostRepository(coll *mongo.Collection) repository.PostRepository {
	return &postRepository{
		coll: coll,
	}
}

func (r *postRepository) Create(post *models.Post) error {
	ctx := context.Background()
	_, err := r.coll.InsertOne(ctx, post)
	return err
}

func (r *postRepository) Update(id primitive.ObjectID, post *models.Post) error {
	ctx := context.Background()
	_, err := r.coll.UpdateOne(ctx, &models.Post{
		ID: id,
	}, post)
	return err
}

func (r *postRepository) FindByID(id primitive.ObjectID) (*models.Post, error) {
	ctx := context.Background()
	result := &models.Post{}

	if err := r.coll.
		FindOne(ctx, &models.Post{
			ID: id,
		}).
		Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
func (r *postRepository) FindBySlug(slug string) (*models.Post, error) {
	ctx := context.Background()
	result := &models.Post{}

	if err := r.coll.
		FindOne(ctx, &models.Post{
			Slug: slug,
		}).
		Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *postRepository) FindAll(offset, limit int64, categoryID primitive.ObjectID) ([]*models.Post, error) {
	ctx := context.Background()
	results := []*models.Post{}
	var otp options.FindOptions
	if offset <= 0 {
		otp = options.FindOptions{
			Limit: &limit,
			Sort: bson.M{
				"created_at": -1,
			},
		}
	} else {
		otp = options.FindOptions{
			Skip:  &offset,
			Limit: &limit,
		}
	}
	filter := bson.M{}
	if !categoryID.IsZero() {
		filter["categoryID"] = categoryID
	}
	cur, err := r.coll.Find(ctx, filter, &otp)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result *models.Post
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

func (r *postRepository) Delete(id primitive.ObjectID) error {
	ctx := context.Background()

	if _, err := r.coll.DeleteOne(ctx, &models.Post{ID: id}); err != nil {
		return err
	}
	return nil
}

func (r *postRepository) Count() (int64, error) {
	ctx := context.Background()
	return r.coll.CountDocuments(ctx, bson.D{})
}
