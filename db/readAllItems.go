package db

import (
	"context"
	"fmt"

	"time"

	"github.com/JeanCntrs/airbnb-catalog-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadAllItems : 2
func ReadAllItems(page int64, search string) ([]*models.Item, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("sample_airbnb")
	coll := db.Collection("listingsAndReviews")

	var results []*models.Item

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 50)
	findOptions.SetLimit(50)
	findOptions.SetProjection(
		bson.M{
			"_id":               1,
			"name":              1,
			"summary":           1,
			"property_type":     1,
			"number_of_reviews": 1,
			"images":            1,
			"address":           1,
			"reviews":           1,
		},
	)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	/* query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
		"$or": []bson.M{
			{"summary": bson.M{"$regex": `(?i)` + search}},
			{"description": bson.M{"$regex": `(?i)` + search}}},
	} */

	/* query := bson.M{"$or": []bson.M{{"name": bson.M{"$regex": `(?i)` + search}, "summary": bson.M{"$regex": `(?i)` + search}}}} */

	/* query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
		"$or": []interface{}{
			bson.M{"summary": bson.M{"$regex": `(?i)` + search}},
		},
	} */

	documents, docError := coll.CountDocuments(ctx, query)
	if docError != nil {
		fmt.Println(docError.Error())
		return results, false
	}

	cursor, err := coll.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	for cursor.Next(ctx) {
		var item models.Item

		err := cursor.Decode(&item)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		item.Documents = documents

		if len(item.Reviews) > 0 {
			item.Reviews = item.Reviews[len(item.Reviews)-1:]
		}

		results = append(results, &item)
	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cursor.Close(ctx)

	return results, true
}
