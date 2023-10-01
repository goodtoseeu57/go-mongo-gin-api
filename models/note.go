package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	UserId      string             `bson:"user_id,omitempty"`
}
