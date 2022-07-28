package userrepository

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (u *UserRepository) AddUser(ctx context.Context, user *entities.User) (*entities.User, *errors.CustomError) {
	
	newUser, err := u.Model.AddUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return newUser, nil

}
