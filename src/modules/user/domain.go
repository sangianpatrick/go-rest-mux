package user

import (
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
)

// Domain contains user's behavior.
type Domain interface {
	GetByID(id string) *wrapper.Property
	GetByEmail(email string) *wrapper.Property
	GetAll() *wrapper.Property
}
