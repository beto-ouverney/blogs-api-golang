package blogpostcontroller

import (
	"context"
	"encoding/json"

	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (c *BlogPostController) GetAll(ctx context.Context) ([]byte, *errors.CustomError) {
	blogPosts, err := c.UseCase.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	bJson, errJ := json.MarshalIndent(&blogPosts, "", "  ")
	if errJ != nil {
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostblogpost.GetAllBlogPosts", Err: errJ}
	}
	return bJson, nil
}
