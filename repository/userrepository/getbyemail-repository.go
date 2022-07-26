package userrepository

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (*entities.User, *errors.CustomError) {
	user, err := u.Model.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
