package usermodel

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (model *modelSqlx) GetByID(ctx context.Context, id int64) (*entities.UserWithoutPassword, *errors.CustomError) {
	var user entities.UserWithoutPassword

	err := model.sqlx.GetContext(ctx, &user, `
		SELECT 
			id,
			displayName,
			email,
			image
		FROM Users 
		WHERE id=?
	`, id)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "usermodel.GetByEmail", Err: err}
	}
	return &user, nil

}
