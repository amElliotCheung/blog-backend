package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Content     string             `json:"content" bson:"content"`
	ReleaseDate time.Time          `json:"release_date,omitempty" bson:"release_date"`
	Comments    []Comment          `json:"comments" bson:"comments"`
}
