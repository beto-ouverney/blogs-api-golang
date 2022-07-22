package usercontroller

import (
	"context"
	"encoding/json"

	"github.com/beto-ouverney/blogs-api-golang/errors"
	"github.com/beto-ouverney/blogs-api-golang/usecase/userusecase"
)

//UserLogin is a controller responsible for transform token string in Json
func (uc *UserController) LoginUser(ctx context.Context, email, password string) (tokenJson []byte, err *errors.CustomError) {
	userUseCase := userusecase.New()
	token, err := userUseCase.LoginUser(ctx, email, password)
	if err == nil {
		data := struct {
			token string
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
