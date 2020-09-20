package db

import (
	"context"
	"fmt"
	"time"

	"github.com/JeanCntrs/airbnb-catalog-server/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ReadItem : Search an item by id in the database
func ReadItem(ID string) (models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConnection.Database("sample_airbnb")
	coll := db.Collection("listingsAndReviews")

	var item models.Item

	condition := bson.M{
		"_id": ID,
	}

	err := coll.FindOne(ctx, condition).Decode(&item)
	if err != nil {
		fmt.Println("Item not found " + err.Error())
		return item, err
	}

	return item, nil
}
