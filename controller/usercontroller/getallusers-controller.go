package usercontroller

import (
	"context"
	"encoding/json"

	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (c *UserController) GetAllUsers(ctx context.Context) ([]byte, *errors.CustomError) {
	users, err := c.UseCase.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	usersJson, errJ := json.MarshalIndent(users, "", "    ")
	if errJ != nil {
		return nil, &errors.CustomError{Op: "usercontroller.GetAllUsers", Err: errJ, Code: errors.EINTERNAL}
	}
	return usersJson, nil
}
