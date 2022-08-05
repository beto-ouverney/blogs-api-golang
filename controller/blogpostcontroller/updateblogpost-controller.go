package blogpostcontroller

import (
	"context"
	"encoding/json"

	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (c BlogPostController) Update(ctx context.Context, id int64, token string, title string, content string) ([]byte, *errors.CustomError) {
	newBlogPost, err := c.UseCase.Update(ctx, id, token, title, content)
	if err != nil {
		return nil, err
	}

	bJson, errJ := json.MarshalIndent(newBlogPost, "", "  ")
	if errJ != nil {
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostblogpost.Update", Err: errJ}
	}
	return bJson, nil
}
