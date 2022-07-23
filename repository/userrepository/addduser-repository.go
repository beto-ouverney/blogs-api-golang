package userrepository

import (
	"context"
	"fmt"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (u *UserRepository) AddUser(ctx context.Context, user *entities.User) (*entities.User, *errors.CustomError) {
	userExists, err := u.GetByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	fmt.Println(userExists)

	newUser, err := u.Model.AddUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return newUser, nil

}
