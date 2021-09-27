package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name,omitempty"`
	Surname  string             `bson:"surname" json:"surname,omitempty"`
	Birth    time.Time          `bson:"birth" json:"birth,omitempty"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password,omitempty"`
	Avatar   string             `bson:"avatar" json:"avatar,omitempty"`
	Banner   string             `bson:"banner" json:"banner,omitempty"`
	Bio      string             `bson:"bio" json:"bio,omitempty"`
	Ubi      string             `bson:"ubi" json:"ubi,omitempty"`
	Web      string             `bson:"web" json:"web,omitempty"`
}
