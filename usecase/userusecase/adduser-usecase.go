package userusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (u *UserUseCase) Add(ctx context.Context, user *entities.User) (*entities.User, *errors.CustomError) {
	userExist, err := u.Repo.GetByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if userExist != nil {
		return nil, &errors.CustomError{Code: errors.ECONFLICT, Op: "userusecase.GetByEmail"}
	}
	newUser, err := u.Repo.Add(ctx, user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
