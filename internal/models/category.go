package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string             `bson:"name,omitempty" json:"name,omitempty"`
	Slug string             `bson:"slug,omitempty"`

	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
}
