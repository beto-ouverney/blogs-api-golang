package categorymodel

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/jmoiron/sqlx"
)

type ICategoryModel interface {
	AddCategory(ctx context.Context, category *entities.Category) (*entities.Category, *errors.CustomError)
	GetAllCategories(ctx context.Context) (*[]entities.Category, *errors.CustomError)
	GetByNameCategory(ctx context.Context, name string) (*entities.Category, *errors.CustomError)
}

type modelSqlx struct {
	sqlx *sqlx.DB
}

func NewSqlxModel(sqlx *sqlx.DB) ICategoryModel {
	return &modelSqlx{sqlx}
}
