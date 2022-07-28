package categoryusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (c *CategoryUseCase) AddCategory(ctx context.Context, category *entities.Category) (*entities.Category, *errors.CustomError) {
	categoryExist, err := c.Repo.GetByNameCategory(ctx, category.Name)
	if categoryExist != nil {
		return nil, &errors.CustomError{Code: errors.ECONFLICT, Op: "categorycategory.AddCategory", Err: err}
	}

	newCategory, err := c.Repo.AddCategory(ctx, category)
	if err == nil {
		return newCategory, nil
	}

	return nil, err
}
