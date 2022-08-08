package usercontroller

import (
	"context"
	"encoding/json"

	"github.com/beto-ouverney/blogs-api-golang/errors"
)

//UserLogin is a controller responsible for transform token string in Json
func (c *UserController) Login(ctx context.Context, email, password string) (tokenJson []byte, err *errors.CustomError) {
	token, err := c.UseCase.Login(ctx, email, password)
	if err == nil {
		data := struct {
			Token string `json:"token"`
		}{
			token,
		}
		tokenJson, errs := json.Marshal(data)

		if errs == nil {
			return tokenJson, nil
		}
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "usercontroller.UserLogin", Err: errs}
	}

	return
}
