package validation

import (
	"net/http"
	"regexp"

	"gitlab.com/patricksangian/go-rest-mux/helpers/wrapper"
	"gitlab.com/patricksangian/go-rest-mux/src/modules/user/model"
	validator "gopkg.in/go-playground/validator.v9"
)

// IsValidUserRegistrationPayload will validate incoming request payload for user registration
func IsValidUserRegistrationPayload(payload *model.User) *wrapper.Property {
	validate := validator.New()
	validate.RegisterValidation("ina-mobile-phone", inaMobilePhone)
	validate.RegisterValidation("app-password", appPassword)

	err := validate.Struct(payload)
	if err != nil {
		return wrapper.Error(http.StatusBadRequest, err.Error())
	}
	return wrapper.Data(http.StatusOK, nil, "is valid payload")
}

func inaMobilePhone(fl validator.FieldLevel) bool {
	phoneRgx := regexp.MustCompile(`^(628)[0-9]+$`)
	return phoneRgx.MatchString(fl.Field().String()) && len(fl.Field().String()) >= 12
}

func appPassword(fl validator.FieldLevel) bool {
	passwordRgx := regexp.MustCompile("^[a-zA-Z0-9]+$")
	return passwordRgx.MatchString(fl.Field().String()) && len(fl.Field().String()) >= 8
}
