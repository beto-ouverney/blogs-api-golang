package blogpostcontroller

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/beto-ouverney/blogs-api-golang/usecase/blogpostusecase"
)

type IBlogPostController interface {
	Add(ctx context.Context, token string, blogPost *entities.BlogPost) ([]byte, *errors.CustomError)
	GetAll(ctx context.Context) ([]byte, *errors.CustomError)
	GetByID(ctx context.Context, id int64) ([]byte, *errors.CustomError)
	Update(ctx context.Context, id int64, token string, title string, content string) ([]byte, *errors.CustomError)
	Delete(ctx context.Context, id int64, token string) *errors.CustomError
}

type BlogPostController struct {
	UseCase blogpostusecase.IBlogPostUseCase
}

func New() *BlogPostController {
	return &BlogPostController{
		UseCase: blogpostusecase.New(),
	}
}
