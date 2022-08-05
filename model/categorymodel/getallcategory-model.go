package categorymodel

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (model *modelSqlx) GetAll(ctx context.Context) (*[]entities.Category, *errors.CustomError) {

	var categories []entities.Category

	err := model.sqlx.SelectContext(ctx, &categories, `SELECT id AS "categories.id",name AS "categories.name"  FROM Categories`)
	if err != nil {
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "categorymodel.GetAllCategories", Err: err}
	}
	return &categories, nil
}
