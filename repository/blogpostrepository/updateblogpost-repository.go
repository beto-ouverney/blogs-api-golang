package blogpostrepository

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (r *BlogPostRepository) Update(ctx context.Context, blogPost *entities.BlogPost) *errors.CustomError {
	err := r.Model.Update(ctx, blogPost)
	if err != nil {
		return err
	}
	return nil
}
