package tokenjwt

import (
	"log"

	"github.com/beto-ouverney/blogs-api-golang/config"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/golang-jwt/jwt/v4"
)

func ExtractData(tokenStr string) (jwt.MapClaims, *errors.CustomError) {
	hmacSecretString := config.JWT_SECRET
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "tokenjwt.ExtractData", Err: err}
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		log.Printf("Invalid JWT Token")
		return nil, &errors.CustomError{Code: errors.EINVALID, Op: "tokenjwt.ExtractData", Err: err}
	}
}
