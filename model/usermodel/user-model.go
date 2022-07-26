package usermodel

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/jmoiron/sqlx"
)

type IUserModel interface {
	GetByEmail(ctx context.Context, email string) (*entities.User, *errors.CustomError)
	AddUser(ctx context.Context, user *entities.User) (*entities.User, *errors.CustomError)
	GetAllUsers(ctx context.Context) ([]entities.UserWithoutPassword, *errors.CustomError)
	GetByID(ctx context.Context, id int64) (*entities.UserWithoutPassword, *errors.CustomError)
}

type modelSqlx struct {
	sqlx *sqlx.DB
}

func NewSqlxModel(sqlx *sqlx.DB) IUserModel {
	return &modelSqlx{sqlx}
}
