package framework

import (
  "errors"
  "net/http"
)

type Engine struct {
  Router *Router
}

func NewEngine() *Engine {
  return &Engine{
    Router: &Router{},
  }
}

type Router struct {
  routingTable map[string]func(rw http.ResponseWriter, r *http.Request)
}

func (r *Router) Get(pathname string, handler func(rw http.ResponseWriter, r *http.Request)) error {

  if r.routingTable == nil {
    r.routingTable = make(map[string]func(rw http.ResponseWriter, r *http.Request))
  }

  if r.routingTable[pathname] != nil {
    return errors.New("existed")
  }
  r.routingTable[pathname] = handler
  return nil // エラーなし
}

func (h *Engine) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

  if r.Method == "GET" {
    handler := h.Router.routingTable[r.URL.Path]
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
