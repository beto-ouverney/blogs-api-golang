package categoryrepository

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (r *CategoryRepository) GetAll(ctx context.Context) (*[]entities.Category, *errors.CustomError) {
	categories, err := r.Model.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
