package handler

import (
	"encoding/json"
	"net/http"

	"github.com/beto-ouverney/blogs-api-golang/controller/blogpostcontroller"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func AddBlogPost(w http.ResponseWriter, r *http.Request) {
	status := 500
	response := []byte("{\"message\":\"Error\"}")
	defer r.Body.Close()

	token := r.Header.Get("Authorization")

	data := struct {
		Title       string  `json:"title"`
		Content     string  `json:"content"`
		CategoryIDs []int64 `json:"categoryIds"`
	}{}

	errJ := json.NewDecoder(r.Body).Decode(&data)
	if errJ != nil {
		errorReturn(w, r, 500, errJ.Error())
	}

	controller := blogpostcontroller.New()
	newBlogPost, errC := controller.Add(r.Context(), token, data.Title, data.Content, data.CategoryIDs)

	if errC != nil {
		if errC.Code != errors.ECONFLICT {
			errorReturn(w, r, 500, errC.Error())
		}
		status = 400
		response = []byte("{\"message\":\"" + errC.Message + "\"}")
	} else {
		status = 201
		response = newBlogPost
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, errW := w.Write(response)
	if errW != nil {
		errorReturn(w, r, 500, errW.Error())
	}

}
