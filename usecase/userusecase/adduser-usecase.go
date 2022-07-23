package userusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (u *UserUseCase) AddUser(ctx context.Context, user *entities.User) (*entities.User, *errors.CustomError) {
	newUser, err := u.Repo.AddUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
