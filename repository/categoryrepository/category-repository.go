package categoryrepository

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/config"
	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/beto-ouverney/blogs-api-golang/model/categorymodel"
	"github.com/jmoiron/sqlx"
)

type ICategoryRepository interface {
	AddCategory(ctx context.Context, category *entities.Category) (*entities.Category, *errors.CustomError)
	GetAllCategories(ctx context.Context) (*[]entities.Category, *errors.CustomError)
	GetByNameCategory(ctx context.Context, name string) (*entities.Category, *errors.CustomError)
	GetByID(ctx context.Context, id int64) (*entities.Category, *errors.CustomError)
}

type CategoryRepository struct {
	Model categorymodel.ICategoryModel
}

type Options struct {
	Sqlx *sqlx.DB
}

// New create a new instance
func New() *CategoryRepository {
	opts := Options{config.GetSqlx()}

	return &CategoryRepository{
		Model: categorymodel.NewSqlxModel(opts.Sqlx),
	}
}
