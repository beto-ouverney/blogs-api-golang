package blogpostcontroller

import (
	"context"
	"encoding/json"

	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (c *BlogPostController) Search(ctx context.Context, search string) ([]byte, *errors.CustomError) {
	blogPosts, errU := c.UseCase.Search(ctx, search)
	if errU != nil {
		return nil, errU
	}

	bJson, err := json.MarshalIndent(blogPosts, "", "  ")
	if err != nil {
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostcontroller.Search", Err: err}
	}

	return bJson, nil
}
