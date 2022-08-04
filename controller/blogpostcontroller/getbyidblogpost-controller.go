package blogpostcontroller

import (
	"context"
	"encoding/json"

	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (c *BlogPostController) GetByIDBlogPost(ctx context.Context, id int64) ([]byte, *errors.CustomError) {
	blogPost, err := c.UseCase.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	bJson, errJ := json.MarshalIndent(&blogPost, "", "  ")
	if errJ != nil {
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostblogpost.GetByID", Err: errJ}
	}
	return bJson, nil
}
