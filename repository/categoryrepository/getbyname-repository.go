package categoryrepository

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (r *CategoryRepository) GetByName(ctx context.Context, name string) (*entities.Category, *errors.CustomError) {
	category, err := r.Model.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return category, nil
}
