package userrepository

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (u *UserRepository) GetByID(ctx context.Context, id int64) (*entities.UserWithoutPassword, *errors.CustomError) {
	user, err := u.Model.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
