package usermodel

import (
	"context"
	"fmt"
	"testing"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/model/usermodel/mock"
	"github.com/golang/mock/gomock"
)

func TestSqlx_GetAllUsers(t *testing.T) {
	ctx := context.Background()
	controller := gomock.NewController(t)
	defer controller.Finish()

	res := gomock.AssignableToTypeOf(context.Background()).Matches(ctx)
	fmt.Println(res) // fasle, but I want true!!

	model := mock.NewMockIUserModel(controller)

	usersMock := []entities.UserWithoutPassword{
		{
			ID:          1,
			DisplayName: "John Doe",
			Email:       "thor@marvel.com",
			Image:       "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50?s=200",
		},
		{
			ID:          2,
			DisplayName: "Jane Doe",
			Email:       "janefoster@marvel.com",
			Image:       "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50?s=200",
		},
	}
	model.EXPECT().GetAllUsers(ctx).Return(usersMock, nil)
}
