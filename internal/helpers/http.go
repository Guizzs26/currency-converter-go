package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func WriteJSON(w http.ResponseWriter, data any, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578 // 1MB

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func WriteErrorJSON(w http.ResponseWriter, message string, status int) error {
	type envolope struct {
		Error string `json:"error"`
	}

	return WriteJSON(w, &envolope{Error: message}, status)
}

func ValidationErrorJSON(w http.ResponseWriter, r *http.Request, verr error) {
	errors := make(map[string]string)

	if ve, ok := verr.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := e.Field()
			tag := e.Tag()

			switch tag {
			case "required":
				errors[field] = "is required"
			case "len":
				errors[field] = "must be exactly " + e.Param() + " characters long"
			case "gt":
				errors[field] = "must be greater than " + e.Param()
			case "uppercase":
				errors[field] = "must be uppercase"
			default:
				errors[field] = "is invalid"
			}
		}
	}

	response := map[string]any{
		"error":  "validation failed",
		"fields": errors,
	}

	WriteJSON(w, response, http.StatusBadRequest)
}
