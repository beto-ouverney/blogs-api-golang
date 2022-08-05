package handler

import (
	"net/http"

	"github.com/beto-ouverney/blogs-api-golang/controller/blogpostcontroller"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func GetAllBlogPosts(w http.ResponseWriter, r *http.Request) {
	status := 500
	response := []byte("{\"message\":\"Error\"}")

	controller := blogpostcontroller.New()

	response, err := controller.GetAll(r.Context())
	if err != nil {
		if err.Code != errors.ECONFLICT {
			errorReturn(w, r, 500, err.Error())
		}
		status = 400
		response = []byte("{\"message\":\"" + err.Message + "\"}")
	} else {
		status = 200
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, errW := w.Write(response)
	if errW != nil {
		errorReturn(w, r, 500, errW.Error())
	}
}
