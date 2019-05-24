package auth

import (
	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/auth/model"
)

// Domain contains Auth behavior
type Domain interface {
	SignIn(payload *model.Auth) *wrapper.Property
}
