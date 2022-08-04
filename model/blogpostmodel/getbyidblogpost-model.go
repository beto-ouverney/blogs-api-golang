package blogpostmodel

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (m *modelSqlx) GetByID(ctx context.Context, id int64) (*entities.BlogPostResponse, *errors.CustomError) {
	var blogPost entities.BlogPostResponse
	err := m.sqlx.GetContext(ctx, &blogPost, `SELECT BlogPosts.id, BlogPosts.title, BlogPosts.content, BlogPosts.published, BlogPosts.updated, user.id AS "user.id", user.displayName AS "user.displayName", user.email AS "user.email", user.image AS "user.image" FROM BlogPosts INNER JOIN Users AS user ON user.id = userId WHERE BlogPosts.id = ?`, id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostmodel.GetByIDBlogPosts", Err: err}
	}
	var categories []entities.Category
	err = m.sqlx.SelectContext(ctx, &categories, `SELECT categories.id AS "categories.id", categories.name AS "categories.name" FROM Categories as categories INNER JOIN PostCategories ON PostCategories.categoryId = categories.id WHERE PostCategories.postId = ?`, blogPost.ID)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostmodel.GetByIDBlogPosts", Err: err}
	}
	blogPost.Categories = categories
	return &blogPost, nil
}
