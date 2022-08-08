package usercontroller

import (
	"context"
	"encoding/json"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (c *UserController) Add(ctx context.Context, dysplayName, email, password, image string) (talkerJson []byte, err *errors.CustomError) {

	user := &entities.User{
		DisplayName: dysplayName,
		Email:       email,
		Password:    password,
		Image:       image,
	}

	newUser, err := c.UseCase.Add(ctx, user)
	if err != nil {
		return nil, err
	}
	newUserJson, errJ := json.MarshalIndent(newUser, "", "    ")
	if errJ == nil {
		return newUserJson, nil
	}
	return nil, &errors.CustomError{Op: "usercontroller.AddUser", Err: errJ, Code: errors.EINTERNAL}
}
