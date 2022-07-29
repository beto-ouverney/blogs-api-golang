package blogpostusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/beto-ouverney/blogs-api-golang/helper/tokenjwt"
)

func verifyFields(blogPost *entities.BlogPost) *errors.CustomError {
	if blogPost.Title == "" {
		return &errors.CustomError{Code: errors.ECONFLICT, Op: "blogpostblogpost.AddBlogPost", Message: errors.ErrorResponse["missingFields"].Message}
	}
	if blogPost.Content == "" {
		return &errors.CustomError{Code: errors.ECONFLICT, Op: "blogpostblogpost.AddBlogPost", Message: errors.ErrorResponse["missingFields"].Message}
	}

	return nil
}

func (u *BlogPostUseCase) AddBlogPost(ctx context.Context, token string, blogPost *entities.BlogPost) (*entities.BlogPost, *errors.CustomError) {
	errF := verifyFields(blogPost)
	if errF != nil {
		return nil, errF
	}

	for _, v := range blogPost.CategoryIDs {
		exits, errC := u.RepoCategory.GetByID(ctx, v)
		if errC != nil {
			return nil, errC
		}
		if exits == nil {
			return nil, &errors.CustomError{Code: errors.ECONFLICT, Op: "blogpostblogpost.AddBlogPost", Message: errors.ErrorResponse["categoryNotFound"].Message}
		}
	}
	claims, errT := tokenjwt.ExtractData(token)
	if errT != nil {
		return nil, errT
	}

	user, errU := u.RepoUser.GetByEmail(ctx, claims["email"].(string))
	if errU != nil {
		return nil, errU
	}
	blogPost.UserID = user.ID

	newBlogPost, err := u.RepoBlogPost.AddBlogPost(ctx, blogPost)

	if err != nil {
		return nil, err
	}

	return newBlogPost, nil
}
