package categoryrepository

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (c *CategoryRepository) Add(ctx context.Context, category *entities.Category) (*entities.Category, *errors.CustomError) {

	newCategory, err := c.Model.Add(ctx, category)

	if err != nil {
		return nil, err
	}

	return newCategory, nil
}
