package userusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (u *UserUseCase) GetAllUsers(ctx context.Context) ([]entities.UserWithoutPassword, *errors.CustomError) {
	return u.Repo.GetAllUsers(ctx)
}
