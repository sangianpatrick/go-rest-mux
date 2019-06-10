package model

import "time"

// User contains app user's property.
type User struct {
	ID        string    `json:"id" bson:"id,omitempty" validate:"-"`
	Name      string    `json:"name" bson:"name,omitempty" validate:"required"`
	Email     string    `json:"email" bson:"email,omitempty" validate:"required,email"`
	Phone     string    `json:"phone" bson:"phone,omitempty" validate:"idn-mobile-number,required"`
	Password  string    `json:"password,omitempty" bson:"password,omitempty" validate:"password,required"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

// Users contains list of user.
type Users []*User

// regexp= for password
