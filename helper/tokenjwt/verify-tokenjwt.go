package tokenjwt

import (
	"time"

	"github.com/beto-ouverney/blogs-api-golang/config"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/golang-jwt/jwt/v4"
)

type claims struct {
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	Image       string `json:"image"`
	jwt.StandardClaims
}

func VerifyToken(token string) (bool, *errors.CustomError) {
	tkn, err := jwt.ParseWithClaims(
		token,
		&claims{},
		func(tk *jwt.Token) (interface{}, error) {
			return []byte(config.JWT_SECRET), nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, nil
		}
		return false, &errors.CustomError{Code: errors.EINTERNAL, Op: "tokenjwt.VerifyToken", Err: err}
	}
	if !tkn.Valid {
		return false, nil
	}

	claims, ok := tkn.Claims.(*claims)
	if !ok {
		return false, nil
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > config.TOKEN_TIME*time.Minute {
		return false, nil
	}
	return true, nil
}
