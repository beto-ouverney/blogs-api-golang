package handler

import (
	"net/http"
	"strconv"

	"github.com/beto-ouverney/blogs-api-golang/controller/blogpostcontroller"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func GetByIDBlogPost(w http.ResponseWriter, r *http.Request) {
	status := 500
	response := []byte("{\"message\":\"Error\"}")

	id, errP := strconv.ParseInt(GetURLParam(r, "id"), 10, 64)
	if errP != nil {
		errorReturn(w, r, 500, errP.Error())
	}
	controller := blogpostcontroller.New()
	blogPost, err := controller.GetByIDBlogPost(r.Context(), id)
	if err == nil {
		status = 200
		response = blogPost
	} else {
		if err.Code != errors.ECONFLICT {
			errorReturn(w, r, 500, err.Error())
		}
		status = errors.ErrorResponse["postNotExist"].Status
		response = []byte("{\"message\":\"" + errors.ErrorResponse["postNotExist"].Message + "\"}")
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, errW := w.Write(response)
	if errW != nil {
		errorReturn(w, r, 500, errW.Error())
	}
}
