package handler

import (
	"net/http"
	"strconv"

	"github.com/beto-ouverney/blogs-api-golang/controller/usercontroller"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func GetByID(w http.ResponseWriter, r *http.Request) {
	status := 500
	response := []byte("{\"message\":\"Error\"}")

	controller := usercontroller.New()
	id, errP := strconv.ParseInt(GetURLParam(r, "id"), 10, 64)
	if errP != nil {
		errorReturn(w, r, 500, errP.Error())
	}
	user, err := controller.GetByID(r.Context(), id)
	if err == nil {
		response = user
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
