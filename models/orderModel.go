package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id"`
	Order_Date time.Time          `json:"order_date" validat: "required"`
	Created_at time.Time          `json:"created_at"`
	Update_at  time.Time          `json:"updated_at"`
	Table__id  *string            `josn:"table_id" validate:"required"`
	Order_id   string             `json"order_id" validate:"required"`
}
