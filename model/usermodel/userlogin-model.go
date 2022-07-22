package usermodel

import (
	"context"
	"fmt"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/jmoiron/sqlx"
)

func NewSqlxModel(sqlx *sqlx.DB) IUserModel {
	return &modelSqlx{sqlx}
}

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
		fmt.Println(err)
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "usermodel.GetByID", Err: err}
	}
	return &user, nil

}
