package categorymodel

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (c *modelSqlx) GetByName(ctx context.Context, name string) (*entities.Category, *errors.CustomError) {
	var category entities.Category
	err := c.sqlx.GetContext(ctx, &category, `SELECT id AS "categories.id" , name AS "categories.name" FROM Categories WHERE name = ?`, name)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "categorymodel.GetByNameCategory", Err: err}
	}
	return &category, nil

}
