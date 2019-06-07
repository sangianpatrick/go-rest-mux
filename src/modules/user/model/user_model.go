package model

// User contains app user's property.
type User struct {
	ID       string `json:"id" bson:"id" validate:"-"`
	Name     string `json:"name" bson:"name" validate:"required"`
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Phone    string `json:"phone" bson:"phone" validate:"idn-mobile-number,required"`
	Password string `json:"password,omitempty" bson:"password" validate:"password,required"`
}

// Users contains list of user.
type Users []*User

// regexp= for password
