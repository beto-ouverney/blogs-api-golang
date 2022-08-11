package handler

import (
	"net/http"

	"github.com/beto-ouverney/blogs-api-golang/controller/blogpostcontroller"
)

func SearchBlogPost(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	defer r.Body.Close()

	controller := blogpostcontroller.New()
	response, err := controller.Search(r.Context(), query)
	if err != nil {
		errorReturn(w, r, 500, err.Error())
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	_, errR := w.Write(response)
	if errR != nil {
		errorReturn(w, r, 500, errR.Error())
	}
}
