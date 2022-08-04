package blogpostusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (u *BlogPostUseCase) GetAllBlogPosts(ctx context.Context) (*[]entities.BlogPostResponse, *errors.CustomError) {
	blogPosts, err := u.RepoBlogPost.GetAllBlogPosts(ctx)
	if err != nil {
		return nil, err
	}
	return blogPosts, nil
}
