package usercontroller

import (
	"context"
	"encoding/json"

	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (c *UserController) GetByID(ctx context.Context, id int64) ([]byte, *errors.CustomError) {

	user, err := c.UseCase.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}
	userJson, errJ := json.MarshalIndent(user, "", "    ")
	if errJ != nil {
		return nil, &errors.CustomError{Op: "usercontroller.GetByID", Err: errJ}
	}
	return userJson, nil
}
