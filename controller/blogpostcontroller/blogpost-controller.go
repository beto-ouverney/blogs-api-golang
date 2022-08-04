package blogpostcontroller

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/beto-ouverney/blogs-api-golang/usecase/blogpostusecase"
)

type IBlogPostController interface {
	AddBlogPost(ctx context.Context, token string, blogPost *entities.BlogPost) ([]byte, *errors.CustomError)
	GetAllBlogPosts(ctx context.Context) ([]byte, *errors.CustomError)
}

type BlogPostController struct {
	UseCase blogpostusecase.IBlogPostUseCase
}

func New() *BlogPostController {
	return &BlogPostController{
		UseCase: blogpostusecase.New(),
	}
}
