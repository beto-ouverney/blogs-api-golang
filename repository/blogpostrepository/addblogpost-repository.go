package blogpostrepository

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (r *BlogPostRepository) AddBlogPost(ctx context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, *errors.CustomError) {

	newBlogPost, err := r.Model.AddBlogPost(ctx, blogPost)

	if err != nil {
		return nil, err
	}

	return newBlogPost, nil
}
