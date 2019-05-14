package user

import (
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user/model"
)

// Domain contains user's behavior.
type Domain interface {
	Create(user *model.User) *wrapper.Property
	GetByEmail(email string) *wrapper.Property
	GetByID(id string) *wrapper.Property
	GetAll(limit int, skip int) *wrapper.Property
}
