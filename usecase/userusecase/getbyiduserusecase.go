package userusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (u *UserUseCase) GetByID(ctx context.Context, id int64) (*entities.UserWithoutPassword, *errors.CustomError) {
	user, err := u.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, &errors.CustomError{Code: errors.ECONFLICT, Op: "userusecase.GetByEmail"}
	}

	return user, nil
}
