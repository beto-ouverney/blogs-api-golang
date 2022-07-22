package middleware

import (
	"encoding/json"
	"regexp"

	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func validateFields(email, password string) (bool, int, string) {
	if email == "" || password == "" {
		return false, errors.ErrorResponse["invalidFields"].Status, errors.ErrorResponse["invalidFields"].Message
	}
	if isEmailValid(email) {
		return false, errors.ErrorResponse["emailIsRequired"].Status, errors.ErrorResponse["emailIsRequired"].Message
	}
	return true, 200, ""
}

//LoginFields validates the login fields
func LoginFieldsValidate(header map[string][]string, decoder *json.Decoder) (ok bool, status int, message string) {
	loginFields := struct {
		email    string `json:"email"`
		password string `json:"password"`
	}{}
	err := decoder.Decode(&loginFields)
	if err != nil {
		return false, 400, "Invalid request"
	}
	ok, status, message = validateFields(loginFields.email, loginFields.password)
	return
}
