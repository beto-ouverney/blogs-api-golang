package usermodel

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (m *modelSqlx) Delete(ctx context.Context, id int64) *errors.CustomError {

	var blogPosts []entities.BlogPostResponse
	err := m.sqlx.SelectContext(ctx, &blogPosts, `SELECT BlogPosts.id,BlogPosts.userId FROM BlogPosts WHERE BlogPosts.userId = ?`, id)
	if err != nil {
		return &errors.CustomError{Code: errors.EINTERNAL, Op: "usermodel.Delete", Err: err}
	}

	tran, err := m.sqlx.BeginTx(ctx, nil)
	if err != nil {
		return &errors.CustomError{Code: errors.EINTERNAL, Op: "usermodel.Delete", Err: err}
	}
	defer tran.Rollback()

	for _, v := range blogPosts {
		_, err = tran.ExecContext(ctx, "DELETE FROM PostCategories WHERE postId = ?", v.ID)
		if err != nil {
			return &errors.CustomError{Code: errors.EINTERNAL, Op: "usermodel.Delete", Err: err}
		}
	}

	_, err = tran.ExecContext(ctx, "DELETE FROM BlogPosts WHERE userId = ?", id)

	if err != nil {
		return &errors.CustomError{Code: errors.EINTERNAL, Op: "usermodel.Delete", Err: err}
	}
	_, err = tran.ExecContext(ctx, "DELETE FROM Users WHERE id = ?", id)
	if err != nil {
		return &errors.CustomError{Code: errors.EINTERNAL, Op: "usermodel.Delete", Err: err}
	}

	if err = tran.Commit(); err != nil {
		return &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostmodel.AddBlogPost", Err: err}
	}

	return nil
}
