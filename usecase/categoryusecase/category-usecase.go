package categoryusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/beto-ouverney/blogs-api-golang/repository/categoryrepository"
)

type ICategoryUseCase interface {
	GetAll(ctx context.Context) (*[]entities.Category, *errors.CustomError)
	Add(ctx context.Context, category *entities.Category) (*entities.Category, *errors.CustomError)
}

type CategoryUseCase struct {
	Repo categoryrepository.ICategoryRepository
}

func New() *CategoryUseCase {
	return &CategoryUseCase{
		Repo: categoryrepository.New(),
	}
}
