package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Slug         string             `bson:"slug,omitempty"`
	CategorySlug string             `bson:"category_slug,omitempty"`
	Title        string             `bson:"title,omitempty"`
	Content      string             `bson:"content,omitempty"`
	CategoryID   primitive.ObjectID `bson:"category_id,omitempty"`
	Thumbnail    string             `bson:"thumbnail,omitempty"`
	Describe     string             `bson:"describe,omitempty"`
	HeartCount   int64              `bson:"heart_count,omitempty"`
	View         int64              `bson:"view,omitempty"`
	CreatedAt    time.Time          `bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `bson:"updated_at,omitempty"`
}
