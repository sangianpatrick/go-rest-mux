package user

import (
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
)

// MongoRepositrory conatains mongodb behavior for package user
type MongoRepositrory interface {
	FindByID(ID string) *wrapper.Property
	FindByEmail(email string) *wrapper.Property
	FindAll(limit int, skip int) *wrapper.Property
}
