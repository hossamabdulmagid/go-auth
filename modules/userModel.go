package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User is the model that governs all notes objects retrieved or inserted into the DB
type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"first_name" validate:"required,min=2,max=100"`
	LastName     *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password     *string            `json:"Password" validate:"required,min=6"`
	Email        *string            `json:"email" validate:"email,required"`
	Phone        *string            `json:"phone" validate:"required"`
	Token        *string            `json:"token"`
	RefreshToken *string            `json:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	UserId       string             `json:"user_id"`
}
