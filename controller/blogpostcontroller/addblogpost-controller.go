package blogpostcontroller

import (
	"context"
	"encoding/json"
	"time"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (c *BlogPostController) Add(ctx context.Context, token string, title string, content string, categoryIds []int64) ([]byte, *errors.CustomError) {
	blogPost := &entities.BlogPost{
		Title:       title,
		Content:     content,
		CategoryIDs: categoryIds,
	}

	newBlogPost, err := c.UseCase.Add(ctx, token, blogPost)

	if err != nil {
		return nil, err
	}

	response := struct {
		ID        int64     `json:"id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		UserID    int64     `json:"userId"`
		Published time.Time `json:"published"`
		Updated   time.Time `json:"updated"`
	}{
		ID:        newBlogPost.ID,
		Title:     newBlogPost.Title,
		Content:   newBlogPost.Content,
		UserID:    newBlogPost.UserID,
		Published: newBlogPost.Published,
		Updated:   newBlogPost.Updated,
	}

	newBlogPostJson, errJ := json.MarshalIndent(&response, "", "  ")
	if errJ != nil {
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostblogpost.AddBlogPost", Err: errJ}
	}

	return newBlogPostJson, nil

}
