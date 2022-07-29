package categoryrepository

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (r *CategoryRepository) GetByID(ctx context.Context, id int64) (*entities.Category, *errors.CustomError) {
	category, err := r.Model.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
