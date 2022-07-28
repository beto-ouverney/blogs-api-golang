package categorycontroller

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/beto-ouverney/blogs-api-golang/usecase/categoryusecase"
)

type ICategoryController interface {
	GetAllCategories(ctx context.Context) ([]byte, *errors.CustomError)
	AddCategory(ctx context.Context, name string) ([]byte, *errors.CustomError)
}

type CategoryController struct {
	UseCase categoryusecase.ICategoryUseCase
}

func New() *CategoryController {
	return &CategoryController{
		UseCase: categoryusecase.New(),
	}
}
