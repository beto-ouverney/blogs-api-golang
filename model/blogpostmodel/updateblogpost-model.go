package blogpostmodel

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (m *modelSqlx) Update(ctx context.Context, blogPost *entities.BlogPost) *errors.CustomError {
	_, err := m.sqlx.NamedExecContext(ctx, `UPDATE BlogPosts SET title=:title, content=:content, published=:published, updated=:updated, userId=:userId WHERE id=:id`, blogPost)
	if err != nil {
		return &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostmodel.Update", Err: err}
	}
	return nil
}
