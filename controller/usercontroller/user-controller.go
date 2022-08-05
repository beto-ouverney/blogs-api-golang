package usercontroller

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/beto-ouverney/blogs-api-golang/usecase/userusecase"
)

type IUserController interface {
	Login(ctx context.Context, email, password string) ([]byte, *errors.CustomError)
	Add(ctx context.Context, dysplayName, email, password, image string) ([]byte, *errors.CustomError)
	GetAll(ctx context.Context) ([]byte, *errors.CustomError)
	GetByID(ctx context.Context, id int64) ([]byte, *errors.CustomError)
}

type UserController struct {
	UseCase userusecase.IUserUseCase
}

func New() *UserController {
	return &UserController{
		UseCase: userusecase.New(),
	}
}
