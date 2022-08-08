package blogpostrepository

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/config"
	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/beto-ouverney/blogs-api-golang/model/blogpostmodel"
	"github.com/jmoiron/sqlx"
)

type IBlogPostRepository interface {
	Add(ctx context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, *errors.CustomError)
	GetAll(ctx context.Context) (*[]entities.BlogPostResponse, *errors.CustomError)
	GetByID(ctx context.Context, id int64) (*entities.BlogPostResponse, *errors.CustomError)
	Update(ctx context.Context, blogPost *entities.BlogPost) *errors.CustomError
	Delete(ctx context.Context, id int64) *errors.CustomError
	Search(ctx context.Context, search string) (*[]entities.BlogPostResponse, *errors.CustomError)
}

type BlogPostRepository struct {
	Model blogpostmodel.IBlogPostModel
}

type Options struct {
	Sqlx *sqlx.DB
}

func New() *BlogPostRepository {
	opts := Options{config.GetSqlx()}

	return &BlogPostRepository{
		Model: blogpostmodel.NewSqlxModel(opts.Sqlx),
	}
}
