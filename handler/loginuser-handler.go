package handler

import (
	"encoding/json"
	"net/http"

	"github.com/beto-ouverney/blogs-api-golang/controller/usercontroller"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	status := 500
	response := []byte("{\"message\":\"Error\"}")

	decoder := json.NewDecoder(r.Body)
	data := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	err := decoder.Decode(&data)
	if err != nil {
		errorReturn(w, r, 500, err.Error())
	}

	defer r.Body.Close()
	controller := usercontroller.New()
	token, errC := controller.LoginUser(r.Context(), data.Email, data.Password)

	if errC != nil {
		if errC.Code == errors.ECONFLICT {
			status = errors.ErrorResponse["invalidFields"].Status
			response = []byte(errors.ErrorResponse["invalidFields"].Message)
		}
		errorReturn(w, r, 500, err.Error())
	}

	status = 200
	response = token

	_, err = w.Write(response)
	if err != nil {
		errorReturn(w, r, 500, err.Error())
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
}
