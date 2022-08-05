package usermodel

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (model *modelSqlx) Add(ctx context.Context, user *entities.User) (*entities.User, *errors.CustomError) {

	result, err := model.sqlx.ExecContext(ctx, `
		INSERT INTO Users (
			displayName,
			email,
			password,
			image
		) VALUES (?, ?, ?, ?)
	`, user.DisplayName, user.Email, user.Password, user.Image)
	if err != nil {
		return nil, &errors.CustomError{Code: errors.EINVALID, Op: "usermodel.AddUser", Err: err}
	}
	user.ID, err = result.LastInsertId()
	if err != nil {
		return nil, &errors.CustomError{Code: errors.EINVALID, Op: "usermodel.AddUser", Err: err}
	}
	return user, nil
}
