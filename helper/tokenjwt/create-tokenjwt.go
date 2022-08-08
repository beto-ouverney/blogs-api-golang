package tokenjwt

import (
	"time"

	"github.com/beto-ouverney/blogs-api-golang/config"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/golang-jwt/jwt/v4"
)

// CreateToken a new token object, specifying signing method and the claims
func CreateToken(displayName string, email string, image string) (string, *errors.CustomError) {
	claims := jwt.MapClaims{}
	claims["displayName"] = displayName
	claims["email"] = email
	claims["image"] = image
	claims["exp"] = time.Now().Add(time.Minute * config.TOKEN_TIME).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.JWT_SECRET))

	if err != nil {
		return "", &errors.CustomError{Code: errors.EINTERNAL, Op: "tokenjwt.CreateToken", Err: err}
	}

	return tokenString, nil
}
