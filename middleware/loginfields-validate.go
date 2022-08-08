package middleware

import (
	"encoding/json"

	"github.com/beto-ouverney/blogs-api-golang/errors"
)

//LoginFields validates the login fields
func LoginFieldsValidate(header map[string][]string, decoder *json.Decoder) (ok bool, status int, message string) {
	fields := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	err := decoder.Decode(&fields)
	if err != nil {
		return false, 400, "Invalid request"
	}
	if fields.Email == "" || fields.Password == "" {
		return false, errors.ErrorResponse["invalidFields"].Status, errors.ErrorResponse["invalidFields"].Message
	}
	if !isEmailValid(fields.Email) {
		return false, errors.ErrorResponse["emailIsRequired"].Status, errors.ErrorResponse["emailIsRequired"].Message
	}
	return true, 200, ""
}
