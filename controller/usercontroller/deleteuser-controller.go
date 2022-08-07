package usercontroller

import (
	"context"

	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (c *UserController) Delete(ctx context.Context, token string) *errors.CustomError {
	return c.UseCase.Delete(ctx, token)
}
