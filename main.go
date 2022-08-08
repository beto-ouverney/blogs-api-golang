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
	router.Route(http.MethodDelete, "/user/me", []myrouter.Middleware{middleware.VerifyToken}, handler.DeleteUser)
	router.Route(http.MethodGet, "/user", []myrouter.Middleware{middleware.VerifyToken}, handler.GetAllUsers)
	router.Route(http.MethodDelete, `/user/(?P<id>\d+)`, []myrouter.Middleware{middleware.VerifyToken}, handler.DeleteUser)
	router.Route(http.MethodGet, "/categories", []myrouter.Middleware{middleware.VerifyToken}, handler.GetAllCategories)
	router.Route(http.MethodPost, "/categories", []myrouter.Middleware{middleware.VerifyToken}, handler.AddCategory)
	router.Route(http.MethodGet, "/post/search", []myrouter.Middleware{middleware.VerifyToken}, handler.SearchBlogPost)
	router.Route(http.MethodPost, "/post", []myrouter.Middleware{middleware.VerifyToken}, handler.AddBlogPost)
	router.Route(http.MethodPut, `/post/(?P<id>\d+)`, []myrouter.Middleware{middleware.VerifyToken}, handler.UpdateBlogPost)
	router.Route(http.MethodGet, `/post/(?P<id>\d+)`, []myrouter.Middleware{middleware.VerifyToken}, handler.GetByIDBlogPost)
	router.Route(http.MethodGet, "/post", []myrouter.Middleware{middleware.VerifyToken}, handler.GetAllBlogPosts)
	router.Route(http.MethodDelete, `/post/(?P<id>\d+)`, []myrouter.Middleware{middleware.VerifyToken}, handler.DeleteBlogPost)
	http.ListenAndServe(":8080", router)
}
