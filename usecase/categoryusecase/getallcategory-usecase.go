package categoryusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (u *CategoryUseCase) GetAllCategories(ctx context.Context) (*[]entities.Category, *errors.CustomError) {
	categories, err := u.Repo.GetAllCategories(ctx)
	if err == nil {
		return categories, nil
	}

	return nil, err
}
