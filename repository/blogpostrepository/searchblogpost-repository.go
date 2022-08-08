package blogpostrepository

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (r *BlogPostRepository) Search(ctx context.Context, search string) (*[]entities.BlogPostResponse, *errors.CustomError) {
	blogPosts, err := r.Model.Search(ctx, search)
	if err != nil {
		return nil, err
	}
	return blogPosts, nil
}
