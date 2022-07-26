package categorycontroller

import (
	"context"
	"encoding/json"

	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (c *CategoryController) GetAll(ctx context.Context) ([]byte, *errors.CustomError) {
	categories, err := c.UseCase.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	categoriesJson, errJ := json.MarshalIndent(categories, "", "    ")
	if errJ != nil {
		return nil, &errors.CustomError{Op: "categorycontroller.GetAllCategories", Err: errJ, Code: errors.EINTERNAL}
	}

	return categoriesJson, nil
}
