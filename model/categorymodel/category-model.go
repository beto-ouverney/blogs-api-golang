package categorymodel

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/jmoiron/sqlx"
)

type ICategoryModel interface {
	Add(ctx context.Context, category *entities.Category) (*entities.Category, *errors.CustomError)
	GetAll(ctx context.Context) (*[]entities.Category, *errors.CustomError)
	GetByName(ctx context.Context, name string) (*entities.Category, *errors.CustomError)
	GetByID(ctx context.Context, id int64) (*entities.Category, *errors.CustomError)
}

type modelSqlx struct {
	sqlx *sqlx.DB
}

func NewSqlxModel(sqlx *sqlx.DB) ICategoryModel {
	return &modelSqlx{sqlx}
}
