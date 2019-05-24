package user

import (
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user/model"
)

// MongoRepositrory conatains mongodb behavior for package user
type MongoRepositrory interface {
	InsertOne(user *model.User) *wrapper.Property
	FindByID(ID string) *wrapper.Property
	FindByEmail(email string) *wrapper.Property
	FindAll(page int, size int) *wrapper.Property
}
