package userusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/errors"
)

//UserInterface is the interface for the user usecase
type IUserUseCase interface {
	LoginUser(ctx context.Context, email, password string) (string, *errors.CustomError)
}

//UserUseCase is the implementation of the user usecase
type UserUseCase struct {
	IUserUseCase
}
