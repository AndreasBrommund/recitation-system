package web

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

type Routes []Route

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.Handler
}

func (this *Routes) Get(name, path string, handlerFunc http.Handler) {
	*this = append(*this, Route{name, "GET", path, handlerFunc})
}

func (this *Routes) Post(name, path string, handlerFunc http.Handler) {
	*this = append(*this, Route{name, "POST", path, handlerFunc})
}

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	for _, route := range initRoutes() {
		router.Handle(route.Method, route.Path, wrapHandler(route.HandlerFunc, route.Name))
	}
	return router
}

func wrapHandler(h http.Handler, name string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		context.Set(r, "name", name)
		h.ServeHTTP(w, r)
	}
}
