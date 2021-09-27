package bd

import (
	"context"
	"log"
	"time"

	"github.com/nativeron/GoTwitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllTweets(ID string, pag int64) ([]*models.ReturnTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("tw")
	col := db.Collection("tweet")

	var result []*models.ReturnTweets

	condition := bson.M{
		"userid": ID,
	}
	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((pag - 1) * 20)

	cursor, err := col.Find(ctx, condition, options)

	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}
	for cursor.Next(context.TODO()) {
		var register models.ReturnTweets
		err := cursor.Decode(&register)
		if err != nil {
			return result, false
		}
		result = append(result, &register)

	}

	return result, true
}
