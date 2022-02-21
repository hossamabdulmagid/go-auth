package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User is the model that governs all notes objects retrieved or inserted into the DB
type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"FirstName" validate:"required,min=2,max=100"`
	LastName     *string            `json:"LastName" validate:"required,min=2,max=100"`
	Password     *string            `json:"Password" validate:"required,min=6"`
	Email        *string            `json:"Email" validate:"email,required"`
	Phone        *string            `json:"Phone" validate:"required"`
	Token        *string            `json:"Token"`
	RefreshToken *string            `json:"RefreshToken"`
	CreatedAt    time.Time          `json:"CreatedAt"`
	UpdatedAt    time.Time          `json:"UpdatedAt"`
	UserId       string             `json:"UserId"`
}
