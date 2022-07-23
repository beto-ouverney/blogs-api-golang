package middleware

import (
	"encoding/json"

	"github.com/beto-ouverney/blogs-api-golang/helper/tokenjwt"

	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func VerifyToken(header map[string][]string, body *json.Decoder) (ok bool, status int, message string) {
	if header["Authorization"] == nil {
		return false, errors.ErrorResponse["tokenNotFound"].Status, errors.ErrorResponse["tokenNotFound"].Message
	}

	token := header["Authorization"][0]
	if len(token) == 0 {
		return false, errors.ErrorResponse["tokenNotFound"].Status, errors.ErrorResponse["tokenNotFound"].Message
	}
	isValid, errT := tokenjwt.VerifyToken(token)
	if errT != nil {
		if errT.Code == errors.EINTERNAL {
			return false, 500, errT.Err.Error()
		}
	}
	if !isValid {
		return false, errors.ErrorResponse["tokenInvalid"].Status, errors.ErrorResponse["tokenInvalid"].Message
	}
	return true, 200, ""
}
