package user

import (
	"github.com/sangianpatrick/go-rest-mux/helpers/wrapper"
	"github.com/sangianpatrick/go-rest-mux/src/modules/user/model"
)

// MongoRepositrory conatains mongodb behavior for package user
type MongoRepositrory interface {
	InsertOne(user *model.User) *wrapper.Property
	FindByID(ID string) *wrapper.Property
	FindByEmail(email string) *wrapper.Property
	FindAll(page int, size int) *wrapper.Property
	UpdateOne(ID string, data *model.User) *wrapper.Property
	DeleteOne(ID string) *wrapper.Property
}
