package validation

import (
	"net/http"

	"github.com/sangianpatrick/go-rest-mux/helpers/wrapper"
	model "github.com/sangianpatrick/go-rest-mux/src/modules/auth/model"
	validator "gopkg.in/go-playground/validator.v9"
)

// IsValidSignInPayload will validate incoming request payload for sign in
func IsValidSignInPayload(payload *model.Auth) *wrapper.Property {
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return wrapper.Error(http.StatusBadRequest, err.Error())
	}
	return wrapper.Data(http.StatusOK, nil, "is valid payload")
}
