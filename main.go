package main

import (
	"net/http"

	"github.com/beto-ouverney/blogs-api-golang/handler"
	"github.com/beto-ouverney/blogs-api-golang/middleware"
	"github.com/beto-ouverney/blogs-api-golang/myrouter"
)

func main() {
	router := &myrouter.Router{}
	router.Route(http.MethodPost, "/login", []myrouter.Middleware{middleware.LoginFieldsValidate}, handler.LoginUser)
	router.Route(http.MethodPost, "/user", []myrouter.Middleware{middleware.AddUserFieldsValidate}, handler.AddUser)
	router.Route(http.MethodGet, "/user", []myrouter.Middleware{middleware.VerifyToken}, handler.GetAllUsers)
	router.Route(http.MethodGet, `/user/(?P<id>\d+)`, []myrouter.Middleware{middleware.VerifyToken}, handler.GetByID)
	router.Route(http.MethodGet, "/categories", []myrouter.Middleware{middleware.VerifyToken}, handler.GetAllCategories)
	router.Route(http.MethodPost, "/categories", []myrouter.Middleware{middleware.VerifyToken}, handler.AddCategory)
	http.ListenAndServe(":8080", router)
}
