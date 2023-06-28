package framework

import (
  "net/http"
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
  // routingTable map[string]func(rw http.ResponseWriter, r *http.Request)
  routingTable TreeNode
}

func (r *Router) Get(pathname string, handler func(rw http.ResponseWriter, r *http.Request)) error {

  r.routingTable.Insert(pathname, handler)
  return nil // エラーなし
}

func (h *Engine) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

  if r.Method == "GET" {
    pathname := r.URL.Path
    handler := h.Router.routingTable.Search(pathname)
    if handler == nil {
      rw.WriteHeader(http.StatusNotFound)
      return
    }
    handler(rw, r)
    return
  }
}

func (e *Engine) Run() {
  http.ListenAndServe("localhost:8080", e)
}
