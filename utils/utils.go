package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New(validator.WithRequiredStructEnabled())

func ParseJson(request *http.Request, payload any) error {
	if request.Body == nil {
		return http.ErrBodyNotAllowed
	}

	return json.NewDecoder(request.Body).Decode(payload)
}

func WriteJson(writer http.ResponseWriter, status int, value any) error {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)

	return json.NewEncoder(writer).Encode(value)
}

func WriteError(writer http.ResponseWriter, status int, err error) {
	WriteJson(writer, status, map[string]string{
		"error": err.Error(),
	})
}
