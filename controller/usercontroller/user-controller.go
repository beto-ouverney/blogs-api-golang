package usercontroller

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/errors"
)

type IUserController interface {
	LoginUser(ctx context.Context, email, password string) ([]byte, *errors.CustomError)
}

type UserController struct {
	IUserController
}

func New() IUserController {
	return &UserController{}
}
