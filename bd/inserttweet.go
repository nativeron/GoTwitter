package bd

import (
	"context"
	"time"

	"github.com/nativeron/GoTwitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTweet(t models.InsertTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("tw")
	col := db.Collection("tweet")

	register := bson.M{
		"userid":  t.UserId,
		"message": t.Message,
		"date":    t.Date,
	}
	result, err := col.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
