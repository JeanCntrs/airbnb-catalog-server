package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type reviews struct {
	DateReview   time.Time `bson:"date" json:"date,omitempty"`
	ReviewerName string    `bson:"reviewer_name" json:"reviewer_name,omitempty"`
	Comments     string    `bson:"comments" json:"comments,omitempty"`
}

// Item : Is the item model
type Item struct {
	ID                   string               `bson:"_id,omitempty" json:"_id"`
	ListingURL           string               `bson:"listing_url" json:"listing_url,omitempty"`
	Name                 string               `bson:"name" json:"name,omitempty"`
	Summary              string               `bson:"summary" json:"summary,omitempty"`
	Space                string               `bson:"space" json:"space,omitempty"`
	Description          string               `bson:"description" json:"description,omitempty"`
	NeighborhoodOverview string               `bson:"neighborhood_overview" json:"neighborhood_overview,omitempty"`
	Notes                string               `bson:"notes" json:"notes,omitempty"`
	Transit              string               `bson:"transit" json:"transit,omitempty"`
	Access               string               `bson:"access" json:"access,omitempty"`
	Interaction          string               `bson:"interaction" json:"interaction,omitempty"`
	HouseRules           string               `bson:"house_rules" json:"house_rules,omitempty"`
	PropertyType         string               `bson:"property_type" json:"property_type,omitempty"`
	RoomType             string               `bson:"room_type" json:"room_type,omitempty"`
	BedType              string               `bson:"bed_type" json:"bed_type,omitempty"`
	MinimumNights        string               `bson:"minimum_nights" json:"minimum_nights,omitempty"`
	MaximumNights        string               `bson:"maximum_nights" json:"maximum_nights,omitempty"`
	CancellationPolicy   string               `bson:"cancellation_policy" json:"cancellation_policy,omitempty"`
	Accommodates         int32                `bson:"accommodates" json:"accommodates,omitempty"`
	Bedrooms             int32                `bson:"bedrooms" json:"bedrooms,omitempty"`
	Beds                 int32                `bson:"beds" json:"beds,omitempty"`
	NumberOfReviews      int32                `bson:"number_of_reviews" json:"number_of_reviews,omitempty"`
	Amenities            []string             `bson:"amenities" json:"amenities,omitempty"`
	Price                primitive.Decimal128 `bson:"price" json:"price,omitempty"`
	Pages                float64              `bson:"pages" json:"pages,omitempty"`
	Images               struct {
		PictureURL string `bson:"picture_url" json:"picture_url,omitempty"`
	}
	Address struct {
		Street  string `bson:"street" json:"street,omitempty"`
		Country string `bson:"country" json:"country,omitempty"`
	}
	Reviews []reviews
}
