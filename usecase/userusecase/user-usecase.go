package userusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/beto-ouverney/blogs-api-golang/repository/userrepository"
)

//UserInterface is the interface for the user usecase
type IUserUseCase interface {
	LoginUser(ctx context.Context, email, password string) (string, *errors.CustomError)
	AddUser(ctx context.Context, user *entities.User) (*entities.User, *errors.CustomError)
}

//UserUseCase is the implementation of the user usecase
type UserUseCase struct {
	Repo userrepository.IUserRepository
}

func New() *UserUseCase {
	return &UserUseCase{
		Repo: userrepository.New(),
	}
}
