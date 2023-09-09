package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string            `json:"last_name" vlaidate: required,min=2,max=100"`
	Email         *string            `json:"email" validate:"email,required"`
	Password      *string            `json:"Passowrd" validate:"required,min=6"`
	Phone         *string            `json: "phone" validate:"required"`
	Token         *string            `json:"token"`
	Avatar        *string            `json:"avatar"`
	Refresh_Token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Update_at     time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}
