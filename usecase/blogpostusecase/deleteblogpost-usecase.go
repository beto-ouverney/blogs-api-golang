package blogpostusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/beto-ouverney/blogs-api-golang/helper/tokenjwt"
)

func (u *BlogPostUseCase) Delete(ctx context.Context, id int64, token string) *errors.CustomError {

	blogPostExist, errB := u.RepoBlogPost.GetByID(ctx, id)
	if errB != nil {
		return errB
	}
	if blogPostExist == nil {
		return &errors.CustomError{Code: errors.ENOTFOUND, Op: "blogpostusecase.Update", Message: errors.ErrorResponse["postNotExist"].Message}
	}

	claims, errT := tokenjwt.ExtractData(token)
	if errT != nil {
		return errT
	}

	user, errU := u.RepoUser.GetByEmail(ctx, claims["email"].(string))
	if errU != nil {
		return errU
	}

	if user.ID != blogPostExist.UserID {
		return &errors.CustomError{Code: errors.ENOTFOUND, Op: "blogpostusecase.Update", Message: errors.ErrorResponse["userNotAuthorized"].Message}
	}

	err := u.RepoBlogPost.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
