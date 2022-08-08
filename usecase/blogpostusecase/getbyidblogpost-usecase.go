package blogpostusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (u *BlogPostUseCase) GetByID(ctx context.Context, id int64) (*entities.BlogPostResponse, *errors.CustomError) {
	blogPost, err := u.RepoBlogPost.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if blogPost == nil {
		return nil, &errors.CustomError{Code: errors.ECONFLICT, Op: "blogpostusecase.GetByID"}
	}
	return blogPost, nil
}
