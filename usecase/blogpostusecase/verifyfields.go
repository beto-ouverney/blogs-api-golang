package blogpostusecase

import (
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func verifyFields(title, content string) *errors.CustomError {
	if title == "" {
		return &errors.CustomError{Code: errors.ECONFLICT, Op: "blogpostblogpost.AddBlogPost", Message: errors.ErrorResponse["missingFields"].Message}
	}
	if content == "" {
		return &errors.CustomError{Code: errors.ECONFLICT, Op: "blogpostblogpost.AddBlogPost", Message: errors.ErrorResponse["missingFields"].Message}
	}

	return nil
}
