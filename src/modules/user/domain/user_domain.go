package domain

import (
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user"
)

// UserDomain contains user properties and use cases
type UserDomain struct {
}

func NewUserDomain() user.Domain {
	return &UserDomain{}
}

func (ud *UserDomain) FindByID(id string) *wrapper.Property
