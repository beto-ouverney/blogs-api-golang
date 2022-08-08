package categorymodel

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (model *modelSqlx) Add(ctx context.Context, category *entities.Category) (*entities.Category, *errors.CustomError) {
	result, err := model.sqlx.ExecContext(ctx, "INSERT INTO Categories (name) VALUES (?)", category.Name)
	if err != nil {

		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "categorymodel.AddCategory", Err: err}
	}

	category.ID, err = result.LastInsertId()
	if err != nil {
		return nil, &errors.CustomError{Code: errors.EINVALID, Op: "usermodel.AddUser", Err: err}
	}

	return category, nil
}
