package handler

import (
	"net/http"
	"strconv"

	"github.com/beto-ouverney/blogs-api-golang/controller/blogpostcontroller"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func DeleteBlogPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id, errP := strconv.ParseInt(GetURLParam(r, "id"), 10, 64)
	if errP != nil {
		errorReturn(w, r, 500, errP.Error())
	}
	token := r.Header.Get("Authorization")
	controller := blogpostcontroller.New()
	err := controller.Delete(r.Context(), id, token)
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
		} else if err.Code == errors.ECONFLICT {
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
