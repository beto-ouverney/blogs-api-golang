package handler

import (
	"net/http"

	"github.com/beto-ouverney/blogs-api-golang/controller/usercontroller"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	status := 500
	response := []byte("{\"message\":\"Error\"}")
	defer r.Body.Close()

	controller := usercontroller.New()
	users, err := controller.GetAll(r.Context())
	if err == nil {
		status = 200
		response = users
	} else {
		if err.Code != errors.ECONFLICT {
			errorReturn(w, r, 500, err.Error())
		}
		status = 500
		response = []byte("{\"message\":\"" + err.Error() + "\"}")
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, errW := w.Write(response)
	if errW != nil {
		errorReturn(w, r, 500, errW.Error())
	}
}
