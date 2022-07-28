package categoryrepository

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (r *CategoryRepository) GetAllCategories(ctx context.Context) (*[]entities.Category, *errors.CustomError) {
	categories, err := r.Model.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
