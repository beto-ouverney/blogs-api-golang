package blogpostmodel

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (m *modelSqlx) Delete(ctx context.Context, id int64) *errors.CustomError {
	tran, err := m.sqlx.BeginTx(ctx, nil)
	if err != nil {
		return &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostmodel.AddBlogPost", Err: err}
	}
	defer tran.Rollback()

	_, err = tran.ExecContext(ctx, `DELETE FROM PostCategories WHERE postId = ?`, id)

	_, err = tran.ExecContext(ctx, `DELETE FROM BlogPosts WHERE id = ?`, id)
	if err != nil {
		return &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostmodel.DeletePost", Err: err}
	}

	if err != nil {
		return &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostmodel.DeletePost", Err: err}
	}

	if err = tran.Commit(); err != nil {
		return &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostmodel.AddBlogPost", Err: err}
	}

	return nil
}
