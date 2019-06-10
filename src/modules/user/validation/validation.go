package validation

import (
	"net/http"
	"regexp"
	"unicode"

	"github.com/sangianpatrick/go-rest-mux/helpers/wrapper"
	"github.com/sangianpatrick/go-rest-mux/src/modules/user/model"
	validator "gopkg.in/go-playground/validator.v9"
)

// IsValidUserRegistrationPayload will validate incoming request payload for user registration
func IsValidUserRegistrationPayload(payload *model.User) *wrapper.Property {
	validate := validator.New()
	validate.RegisterValidation("idn-mobile-number", validIDNMobileNumber)
	validate.RegisterValidation("password", validPassword)

	err := validate.Struct(payload)
	if err != nil {
		return wrapper.Error(http.StatusBadRequest, err.Error())
	}
	return wrapper.Data(http.StatusOK, nil, "is valid payload")
}

func validIDNMobileNumber(fl validator.FieldLevel) bool {
	phoneRgx := regexp.MustCompile(`^(628)[0-9]+$`)
	return phoneRgx.MatchString(fl.Field().String()) && len(fl.Field().String()) >= 12
}

func validPassword(fl validator.FieldLevel) bool {

	// Using regular expression is more efficient, but there's a lack of understanding to RE2 Syntax, that's why i choose this approach.
	// Please feel free to re-code it if you had more efficient way (using RE2 Syntax)
	// Don't ever think that '\A(?>(?<upper>[A-Z])|(?<lower>[a-z])|(?<digit>[0-9])|.){8,}?(?(upper)(?(lower)(?(digit)|(?!))|(?!))|(?!))' or
	// '/^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.{8,})/' could run properly, they are even unsupported in Golang, because of written in Perl Syntax.
	// Just for info, RE2 Syntax is pure regex that has no lookbehind and lookaround feature. You can try this with FSA first. So, happy to them who love Automata Theory :)

	// Don't ever apply this approach in production.
	var (
		passwordString = fl.Field().String()
		hasMinLen      = false
		hasUpper       = false
		hasLower       = false
		hasNumber      = false
		hasSpecial     = false
	)
	if len(passwordString) >= 8 {
		hasMinLen = true
	}
	for _, char := range passwordString {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
