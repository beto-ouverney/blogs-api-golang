package blogpostrepository

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (r *BlogPostRepository) GetByID(ctx context.Context, id int64) (*entities.BlogPostResponse, *errors.CustomError) {
	blogPost, err := r.Model.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return blogPost, nil
}
