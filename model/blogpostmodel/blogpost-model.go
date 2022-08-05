package blogpostmodel

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/jmoiron/sqlx"
)

type IBlogPostModel interface {
	Add(ctx context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, *errors.CustomError)
	GetAll(ctx context.Context) (*[]entities.BlogPostResponse, *errors.CustomError)
	GetByID(ctx context.Context, id int64) (*entities.BlogPostResponse, *errors.CustomError)
	Update(ctx context.Context, blogPost *entities.BlogPost) *errors.CustomError
	Delete(ctx context.Context, id int64) *errors.CustomError
}

type modelSqlx struct {
	sqlx *sqlx.DB
}

func NewSqlxModel(sqlx *sqlx.DB) IBlogPostModel {
	return &modelSqlx{sqlx}
}
