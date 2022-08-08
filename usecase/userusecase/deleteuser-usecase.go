package userusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/beto-ouverney/blogs-api-golang/helper/tokenjwt"
)

func (u *UserUseCase) Delete(ctx context.Context, token string) *errors.CustomError {
	claims, errT := tokenjwt.ExtractData(token)
	if errT != nil {
		return errT
	}
	userExist, errU := u.Repo.GetByEmail(ctx, claims["email"].(string))
	if errU != nil {
		return errU
	}

	if userExist == nil {
		return &errors.CustomError{Code: errors.ENOTFOUND, Op: "userusecase.Delete", Message: errors.ErrorResponse["userNotExist"].Message}
	}

	err := u.Repo.Delete(ctx, userExist.ID)
	if err != nil {
		return err
	}
	return nil
}
