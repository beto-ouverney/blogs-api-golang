package userrepository

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (r *UserRepository) Delete(ctx context.Context, id int64) *errors.CustomError {
	err := r.Model.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
