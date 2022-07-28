package categorycontroller

import (
	"context"
	"encoding/json"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (c *CategoryController) AddCategory(ctx context.Context, name string) ([]byte, *errors.CustomError) {

	category := &entities.Category{Name: name}
	newCategory, err := c.UseCase.AddCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	newCategoryJson, errJ := json.MarshalIndent(newCategory, "", "  ")
	if errJ != nil {
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "categorycontroller.AddCategory", Err: errJ}
	}

	return newCategoryJson, nil
}
