package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	ID      primitive.ObjectID `json:"id" bson:"_id" `
	UserID  primitive.ObjectID `json:"user_id" bson:"user_id"`
	Content string             `json:"content" bson:"content"`
}
