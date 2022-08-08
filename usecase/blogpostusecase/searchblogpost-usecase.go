package blogpostusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (u *BlogPostUseCase) Search(ctx context.Context, search string) (*[]entities.BlogPostResponse, *errors.CustomError) {
	if search == "" {
		blogPosts, err := u.RepoBlogPost.GetAll(ctx)
		if err != nil {
			return nil, err
		}
		return blogPosts, nil
	}
	blogPosts, err := u.RepoBlogPost.Search(ctx, search)
	if err != nil {
		return nil, err
	}
	return blogPosts, nil
}
