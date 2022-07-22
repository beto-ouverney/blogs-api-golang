package main

import (
	"net/http"

	"github.com/beto-ouverney/blogs-api-golang/handler"
	"github.com/beto-ouverney/blogs-api-golang/middleware"
	"github.com/beto-ouverney/blogs-api-golang/myrouter"
)

func main() {
	router := &myrouter.Router{}
	router.Route(http.MethodPost, `/login`, []myrouter.Middleware{middleware.LoginFieldsValidate}, handler.LoginUser)
}
