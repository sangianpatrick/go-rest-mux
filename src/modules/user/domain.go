package user

import (
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
)

// Domain contains user's behavior.
type Domain interface {
	GetProfile(email string) *wrapper.Property
	GetByID(id string) *wrapper.Property
	GetAll(limit int, skip int) *wrapper.Property
}
