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
	AddBlogPost(ctx context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, *errors.CustomError)
	GetAllBlogPosts(ctx context.Context) (*[]entities.BlogPostResponse, *errors.CustomError)
	GetByID(ctx context.Context, id int64) (*entities.BlogPostResponse, *errors.CustomError)
	Update(ctx context.Context, blogPost *entities.BlogPost) *errors.CustomError
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
