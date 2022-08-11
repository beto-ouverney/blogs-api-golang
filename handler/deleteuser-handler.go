package handler

import (
	"net/http"

	"github.com/beto-ouverney/blogs-api-golang/controller/usercontroller"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	controller := usercontroller.New()

	token := r.Header.Get("Authorization")

	err := controller.Delete(r.Context(), token)
	if err == nil {
		w.WriteHeader(204)
		w.Header().Set("Content-Type", "application/json")
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err.Code == errors.ENOTFOUND {
			w.WriteHeader(404)
			_, errW := w.Write([]byte("{\"message\":\"" + err.Message + "\"}"))
			if errW != nil {
				errorReturn(w, r, 500, errW.Error())
			}
		} else {
			errorReturn(w, r, 500, err.Error())
		}
	}
	return
}
