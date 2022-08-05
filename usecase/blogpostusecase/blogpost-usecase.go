package blogpostusecase

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/beto-ouverney/blogs-api-golang/repository/blogpostrepository"
	"github.com/beto-ouverney/blogs-api-golang/repository/categoryrepository"
	"github.com/beto-ouverney/blogs-api-golang/repository/userrepository"
)

type IBlogPostUseCase interface {
	Add(ctx context.Context, token string, blogPost *entities.BlogPost) (*entities.BlogPost, *errors.CustomError)
	GetAll(ctx context.Context) (*[]entities.BlogPostResponse, *errors.CustomError)
	GetByID(ctx context.Context, id int64) (*entities.BlogPostResponse, *errors.CustomError)
	Update(ctx context.Context, id int64, token string, title, content string) (*entities.BlogPostResponse, *errors.CustomError)
	Delete(ctx context.Context, id int64, token string) *errors.CustomError
}

type BlogPostUseCase struct {
	RepoBlogPost blogpostrepository.IBlogPostRepository
	RepoUser     userrepository.IUserRepository
	RepoCategory categoryrepository.ICategoryRepository
}

func New() IBlogPostUseCase {
	return &BlogPostUseCase{
		RepoBlogPost: blogpostrepository.New(),
		RepoUser:     userrepository.New(),
		RepoCategory: categoryrepository.New(),
	}
}
