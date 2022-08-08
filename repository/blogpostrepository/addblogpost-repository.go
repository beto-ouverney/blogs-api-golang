package blogpostrepository

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (r *BlogPostRepository) Add(ctx context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, *errors.CustomError) {

	newBlogPost, err := r.Model.Add(ctx, blogPost)

	if err != nil {
		return nil, err
	}

	return newBlogPost, nil
}
