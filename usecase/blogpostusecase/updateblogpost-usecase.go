package blogpostusecase

import (
	"context"
	"time"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/beto-ouverney/blogs-api-golang/helper/tokenjwt"
)

func (u *BlogPostUseCase) Update(ctx context.Context, id int64, token, title, content string) (*entities.BlogPostResponse, *errors.CustomError) {

	errF := verifyFields(title, content)
	if errF != nil {
		return nil, errF
	}

	blogPostExist, errB := u.RepoBlogPost.GetByID(ctx, id)
	if errB != nil {
		return nil, errB
	}
	if blogPostExist == nil {
		return nil, &errors.CustomError{Code: errors.ENOTFOUND, Op: "blogpostusecase.Update", Message: errors.ErrorResponse["postNotExist"].Message}
	}

	newBlogPost := entities.BlogPost{
		ID:        id,
		Title:     title,
		Content:   content,
		Updated:   time.Now(),
		Published: blogPostExist.Published,
		UserID:    blogPostExist.UserID,
	}

	claims, errT := tokenjwt.ExtractData(token)
	if errT != nil {
		return nil, errT
	}

	user, errU := u.RepoUser.GetByEmail(ctx, claims["email"].(string))
	if errU != nil {
		return nil, errU
	}

	if user.ID != blogPostExist.UserID {
		return nil, &errors.CustomError{Code: errors.ENOTFOUND, Op: "blogpostusecase.Update", Message: errors.ErrorResponse["userNotAuthorized"].Message}
	}

	err := u.RepoBlogPost.Update(ctx, &newBlogPost)
	if err != nil {
		return nil, err
	}

	newBlogPostResponse := entities.BlogPostResponse{
		ID:         id,
		Title:      title,
		Content:    content,
		Updated:    time.Now(),
		Published:  blogPostExist.Published,
		UserID:     blogPostExist.UserID,
		User:       blogPostExist.User,
		Categories: blogPostExist.Categories,
	}
	return &newBlogPostResponse, nil
}
