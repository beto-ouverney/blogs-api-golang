package userusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/config"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/beto-ouverney/blogs-api-golang/helper/tokenjwt"
	"github.com/beto-ouverney/blogs-api-golang/repository/userrepository"
)

func New() IUserUseCase {
	return &UserUseCase{}
}

func (usecase *UserUseCase) LoginUser(ctx context.Context, email, password string) (token string, err *errors.CustomError) {
	repo := userrepository.New(userrepository.Options{
		Sqlx: config.GetSqlx(),
	})

	user, err := repo.GetByEmail(ctx, email)
	if err != nil {
		return token, &errors.CustomError{Op: "userusecase.GetByID", Err: err}
	}

	if user == nil {
		return token, &errors.CustomError{Code: errors.ECONFLICT, Op: "userusecase.GetByID"}
	}

	if user.Password != password {
		return token, &errors.CustomError{Code: errors.ECONFLICT, Op: "userusecase.GetByID"}
	}

	token, err = tokenjwt.CreateToken(user.DisplayName, user.Email, user.Image)
	if err != nil {
		return token, &errors.CustomError{Code: errors.EINTERNAL, Op: "userusecase.GetByID", Err: err}
	}
	return token, nil
}
