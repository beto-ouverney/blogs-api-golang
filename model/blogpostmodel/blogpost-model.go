package blogpostmodel

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/jmoiron/sqlx"
)

type IBlogPostModel interface {
	AddBlogPost(ctx context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, *errors.CustomError)
	GetAllBlogPosts(ctx context.Context) (*[]entities.BlogPostResponse, *errors.CustomError)
	GetByID(ctx context.Context, id int64) (*entities.BlogPostResponse, *errors.CustomError)
}

type modelSqlx struct {
	sqlx *sqlx.DB
}

func NewSqlxModel(sqlx *sqlx.DB) IBlogPostModel {
	return &modelSqlx{sqlx}
}
