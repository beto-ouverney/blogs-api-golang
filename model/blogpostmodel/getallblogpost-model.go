package blogpostmodel

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (m *modelSqlx) GetAllBlogPosts(ctx context.Context) (*[]entities.BlogPostResponse, *errors.CustomError) {
	var blogPosts []entities.BlogPostResponse
	err := m.sqlx.SelectContext(ctx, &blogPosts, `SELECT BlogPosts.id, BlogPosts.title, BlogPosts.content, BlogPosts.published, BlogPosts.updated, user.id AS "user.id", user.displayName AS "user.displayName", user.email AS "user.email", user.image AS "user.image" FROM BlogPosts INNER JOIN Users AS user ON user.id = userId`)
	if err != nil {
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostmodel.GetAllBlogPosts", Err: err}
	}
	var allBlogPosts []entities.BlogPostResponse
	for _, v := range blogPosts {
		var categories []entities.Category
		err = m.sqlx.SelectContext(ctx, &categories, `SELECT categories.id AS "categories.id", categories.name AS "categories.name" FROM Categories as categories INNER JOIN PostCategories ON PostCategories.categoryId = categories.id WHERE PostCategories.postId = ?`, v.ID)

		if err != nil {
			return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostmodel.GetAllBlogPosts", Err: err}
		}
		v.Categories = categories
		allBlogPosts = append(allBlogPosts, v)
	}

	return &allBlogPosts, nil
}
