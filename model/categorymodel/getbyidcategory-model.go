package categorymodel

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (m *modelSqlx) GetByID(ctx context.Context, id int64) (*entities.Category, *errors.CustomError) {
	var category entities.Category
	err := m.sqlx.GetContext(ctx, &category, "SELECT id,name FROM Categories WHERE id = ?", id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "categorymodel.GetByID", Err: err}
	}
	return &category, nil
}
