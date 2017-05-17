package routes

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const params = "params"

//httprouter wrapper
type RouterWrapper struct {
	Router *httprouter.Router
}

//initialization
func NewRouter() *RouterWrapper {
	return &RouterWrapper{Router: httprouter.New()}
}

// Get is a shortcut for router.Handle("GET", path, handle)
func (r *RouterWrapper) Get(path string, fn http.HandlerFunc) {
	r.Router.GET(path, HandlerFunc(fn))
}

// Post is a shortcut for router.Handle("POST", path, handle)
func (r *RouterWrapper) Post(path string, fn http.HandlerFunc) {
	r.Router.POST(path, HandlerFunc(fn))
}
func HandlerFunc(fn http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := context.WithValue(r.Context(), params, p)
		fn(w, r.WithContext(ctx))
	}
}

// Context returns the URL parameters
func Params(r *http.Request) httprouter.Params {
	return r.Context().Value(params).(httprouter.Params)
}
