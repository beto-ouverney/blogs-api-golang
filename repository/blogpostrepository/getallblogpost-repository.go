package blogpostrepository

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (r *BlogPostRepository) GetAll(ctx context.Context) (*[]entities.BlogPostResponse, *errors.CustomError) {
	blogPosts, err := r.Model.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return blogPosts, nil
}
