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
	AddBlogPost(ctx context.Context, token string, blogPost *entities.BlogPost) (*entities.BlogPost, *errors.CustomError)
	GetAllBlogPosts(ctx context.Context) (*[]entities.BlogPostResponse, *errors.CustomError)
	GetByID(ctx context.Context, id int64) (*entities.BlogPostResponse, *errors.CustomError)
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
