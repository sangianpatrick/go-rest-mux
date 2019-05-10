package wrapper

import (
	"encoding/json"
	"net/http"
)

// Property contains properties of response
type Property struct {
	Error   bool        `json:"error"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

// Error returns error response properties.
func Error(code int, message string) Property {
	return Property{
		Error:   true,
		Message: message,
		Code:    code,
	}
}

// Data returns data response properties.
func Data(code int, data interface{}, message string) Property {
	return Property{
		Error:   false,
		Data:    data,
		Message: message,
		Code:    code,
	}
}

// Response returns json data via http.
func Response(w http.ResponseWriter, code int, data *Property, message string) error {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
	return nil
}
