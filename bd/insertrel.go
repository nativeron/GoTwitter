package bd

import (
	"context"
	"time"

	"github.com/nativeron/GoTwitter/models"
)

func InsertRel(t models.Rel) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("tw")
	col := db.Collection("rel")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
