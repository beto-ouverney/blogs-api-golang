package usermodel

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (model *modelSqlx) GetAllUsers(ctx context.Context) ([]entities.UserWithoutPassword, *errors.CustomError) {
	var users []entities.UserWithoutPassword

	err := model.sqlx.SelectContext(ctx, &users, "SELECT id, displayName, email, image FROM Users")

	if err != nil {
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "usermodel.GetAllUsers", Err: err}
	}
	return users, nil
}
