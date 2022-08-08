package blogpostcontroller

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (c *BlogPostController) Delete(ctx context.Context, id int64, token string) *errors.CustomError {
	err := c.UseCase.Delete(ctx, id, token)
	if err != nil {
		return err
	}

	return nil

}
