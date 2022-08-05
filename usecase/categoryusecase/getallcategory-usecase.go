package categoryusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (u *CategoryUseCase) GetAll(ctx context.Context) (*[]entities.Category, *errors.CustomError) {
	categories, err := u.Repo.GetAll(ctx)
	if err == nil {
		return categories, nil
	}

	return nil, err
}
