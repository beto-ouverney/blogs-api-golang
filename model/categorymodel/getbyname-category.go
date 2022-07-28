package categorymodel

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (c *modelSqlx) GetByNameCategory(ctx context.Context, name string) (*entities.Category, *errors.CustomError) {
	var category entities.Category
	err := c.sqlx.GetContext(ctx, &category, "SELECT id,name FROM Categories WHERE name = ?", name)
	if err != nil {
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "categorymodel.GetByNameCategory", Err: err}
	}
	return &category, nil

}
