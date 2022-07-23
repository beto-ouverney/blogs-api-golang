package usercontroller

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/beto-ouverney/blogs-api-golang/usecase/userusecase"
)

type IUserController interface {
	LoginUser(ctx context.Context, email, password string) ([]byte, *errors.CustomError)
	AddUser(ctx context.Context, dysplayName, email, password, image string) ([]byte, *errors.CustomError)
}

type UserController struct {
	UseCase userusecase.IUserUseCase
}

func New() *UserController {
	return &UserController{
		UseCase: userusecase.New(),
	}
}
