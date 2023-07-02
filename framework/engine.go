package framework

import (
	"net/http"
	"strings"
)

type Engine struct {
	Router *Router
}

func NewEngine() *Engine {
	return &Engine{
		Router: &Router{
			routingTable: Contructor(),
		},
	}
}

type Router struct {
	routingTable TreeNode
}

func (r *Router) Get(pathname string, handler func(ctx *MyContext)) error {
	pathname = strings.TrimSuffix(pathname, "/")
	existedHandler := r.routingTable.Search(pathname)

	if existedHandler != nil {
		panic("already exists")
	}

	r.routingTable.Insert(pathname, handler)
	return nil
}

func (h *Engine) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	ctx := NewMyContext(rw, r)

	if r.Method == "GET" {
		pathname := r.URL.Path
		pathname = strings.TrimSuffix(pathname, "/")
		handler := h.Router.routingTable.Search(pathname)

		if handler == nil {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		handler(ctx)
		return
	}
}

func (e *Engine) Run() {
	http.ListenAndServe("localhost:8080", e)
}
