package userrepository

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (u *UserRepository) GetAllUsers(ctx context.Context) ([]entities.UserWithoutPassword, *errors.CustomError) {
	return u.Model.GetAllUsers(ctx)
}
