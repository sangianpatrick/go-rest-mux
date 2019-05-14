package model

// User contains app user's property.
type User struct {
	ID       string `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Phone    string `json:"phone" bson:"phone"`
	Password string `json:"password,omitempty" bson:"password"`
}

// Users contains list of user.
type Users []*User
