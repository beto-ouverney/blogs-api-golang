package handler

import (
	"encoding/json"
	"net/http"

	"github.com/beto-ouverney/blogs-api-golang/controller/categorycontroller"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func AddCategory(w http.ResponseWriter, r *http.Request) {
	status := 500
	response := []byte("{\"message\":\"Error\"}")

	data := struct {
		Name string `json:"name"`
	}{}

	errJ := json.NewDecoder(r.Body).Decode(&data)
	if errJ != nil {
		errorReturn(w, r, 500, errJ.Error())
	}

	controller := categorycontroller.New()
	newCategory, err := controller.Add(r.Context(), data.Name)
	if err != nil {
		if err.Code != errors.ECONFLICT {
			errorReturn(w, r, 500, err.Error())
		}
		status = errors.ErrorResponse["categoryAlreadyExists"].Status
		response = []byte("{\"message\":\"" + errors.ErrorResponse["categoryAlreadyExists"].Message + "\"}")
	} else {
		status = 201
		response = newCategory
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, errW := w.Write(response)
	if errW != nil {
		errorReturn(w, r, 500, errW.Error())
	}

}
