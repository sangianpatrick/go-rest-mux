package wrapper

import (
	"encoding/json"
	"net/http"
)

// Property contains properties of response
type Property struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *meta       `json:"meta,omitempty"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

type meta struct {
	TotalPage       int `json:"totalPage"`
	Page            int `json:"page"`
	TotalData       int `json:"totalData"`
	TotalDataOnPage int `json:"totalDataOnPage"`
}

// Error returns error response properties.
func Error(code int, message string) *Property {
	return &Property{
		Success: false,
		Message: message,
		Code:    code,
	}
}

// Data returns data response properties.
func Data(code int, data interface{}, message string) *Property {
	return &Property{
		Success: true,
		Data:    data,
		Message: message,
		Code:    code,
	}
}

// PaginationData returns data response properties with pagination.
func PaginationData(code int, data interface{}, totalPage int, page int, totalData int, totalDataOnPage int, message string) *Property {
	return &Property{
		Success: true,
		Data:    data,
		Meta: &meta{
			TotalPage:       totalPage,
			Page:            page,
			TotalData:       totalData,
			TotalDataOnPage: totalDataOnPage,
		},
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
