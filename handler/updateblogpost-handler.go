package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/beto-ouverney/blogs-api-golang/controller/blogpostcontroller"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func UpdateBlogPost(w http.ResponseWriter, r *http.Request) {
	status := 500
	response := []byte("{\"message\":\"Error\"}")
	defer r.Body.Close()

	id, errP := strconv.ParseInt(GetURLParam(r, "id"), 10, 64)
	if errP != nil {
		errorReturn(w, r, 500, errP.Error())
	}

	data := struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}{}

	errJ := json.NewDecoder(r.Body).Decode(&data)
	if errJ != nil {
		errorReturn(w, r, 500, errJ.Error())
	}

	token := r.Header.Get("Authorization")

	controller := blogpostcontroller.New()
	response, err := controller.Update(r.Context(), id, token, data.Title, data.Content)
	if err == nil {
		status = 200
	} else if err.Code == errors.ECONFLICT {
		status = 400
		response = []byte("{\"message\":\"" + err.Message + "\"}")
	} else if err.Code == errors.ENOTFOUND {
		status = 404
		fmt.Printf("%+v\n", err)
		response = []byte("{\"message\":\"" + err.Message + "\"}")
	} else {
		errorReturn(w, r, 500, err.Error())
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, errW := w.Write(response)
	if errW != nil {
		errorReturn(w, r, 500, errW.Error())
	}
}
