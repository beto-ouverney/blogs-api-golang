package usermodel

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (model *modelSqlx) GetByEmail(ctx context.Context, email string) (*entities.User, *errors.CustomError) {

	var user entities.User

	err := model.sqlx.GetContext(ctx, &user, `
		SELECT 
			id,
			displayName,
			email,
			password,
			image
		FROM Users 
		WHERE email=?
	`, email)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "usermodel.GetByEmail", Err: err}
	}
	return &user, nil

}
