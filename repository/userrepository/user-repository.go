package userrepository

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/config"
	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/beto-ouverney/blogs-api-golang/model/usermodel"
	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	GetByEmail(ctx context.Context, email string) (*entities.User, *errors.CustomError)
	AddUser(ctx context.Context, user *entities.User) (*entities.User, *errors.CustomError)
}

type UserRepository struct {
	Model usermodel.IUserModel
}

type Options struct {
	Sqlx *sqlx.DB
}

// New create a new instance
func New() *UserRepository {
	opts := Options{config.GetSqlx()}

	return &UserRepository{
		Model: usermodel.NewSqlxModel(opts.Sqlx),
	}
}
