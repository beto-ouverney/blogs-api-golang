package categoryusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (c *CategoryUseCase) Add(ctx context.Context, category *entities.Category) (*entities.Category, *errors.CustomError) {
	categoryExist, err := c.Repo.GetByName(ctx, category.Name)
	if categoryExist != nil {
		return nil, &errors.CustomError{Code: errors.ECONFLICT, Op: "categorycategory.AddCategory", Err: err}
	}

	newCategory, err := c.Repo.Add(ctx, category)
	if err == nil {
		return newCategory, nil
	}

	return nil, err
}
