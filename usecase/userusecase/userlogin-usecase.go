package userusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/beto-ouverney/blogs-api-golang/helper/tokenjwt"
)

func (u *UserUseCase) LoginUser(ctx context.Context, email, password string) (string, *errors.CustomError) {

	user, err := u.Repo.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", &errors.CustomError{Code: errors.ECONFLICT, Op: "userusecase.GetByID"}
	}

	if user.Password != password {
		return "", &errors.CustomError{Code: errors.ECONFLICT, Op: "userusecase.GetByID"}
	}

	token, err := tokenjwt.CreateToken(user.DisplayName, user.Email, user.Image)
	if err != nil {
		return token, err
	}
	return token, nil
}
