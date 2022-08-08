package middleware

import (
	"encoding/json"

	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func AddUserFieldsValidate(header map[string][]string, decoder *json.Decoder) (ok bool, status int, message string) {

	fields := struct {
		DisplayName string `json:"displayName"`
		Image       string `json:"image"`
		Email       string `json:"email"`
		Password    string `json:"password"`
	}{}

	err := decoder.Decode(&fields)
	if err != nil {
		return false, 400, "Invalid request"
	}

	if fields.DisplayName == "" {
		return false, errors.ErrorResponse["displayNameIsRequired"].Status, errors.ErrorResponse["displayNameIsRequired"].Message
	}
	if len(fields.DisplayName) < 8 {
		return false, errors.ErrorResponse["displayNameInvalid"].Status, errors.ErrorResponse["displayNameInvalid"].Message
	}
	if fields.Email == "" {
		return false, errors.ErrorResponse["emailIsRequired"].Status, errors.ErrorResponse["emailIsRequired"].Message
	}
	if !isEmailValid(fields.Email) {
		return false, errors.ErrorResponse["emailIsRequired"].Status, errors.ErrorResponse["emailIsRequired"].Message
	}
	if fields.Password == "" {
		return false, errors.ErrorResponse["passwordIsRequired"].Status, errors.ErrorResponse["passwordIsRequired"].Message
	}
	if len(fields.Password) < 6 {
		return false, errors.ErrorResponse["invalidPassword"].Status, errors.ErrorResponse["invalidPassword"].Message
	}
	return true, 200, ""
}
