package auth

import (
	"github.com/sangianpatrick/go-rest-mux/helpers/wrapper"
	"github.com/sangianpatrick/go-rest-mux/src/modules/auth/model"
)

// Domain contains Auth behavior
type Domain interface {
	SignIn(payload *model.Auth) *wrapper.Property
}
