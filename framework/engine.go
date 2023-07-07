package framework

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

type Engine struct {
	Router *Router
}

func NewEngine() *Engine {
	return &Engine{
		Router: &Router{
			routingTables: map[string]*TreeNode{
				"get":    Contructor(),
				"post":   Contructor(),
				"patch":  Contructor(),
				"put":    Contructor(),
				"delete": Contructor(),
			},
		},
	}
}

type Router struct {
	routingTables map[string]*TreeNode
}

func (r *Router) register(method string, pathname string, handler func(ctx *MyContext)) error {
	routingTable := r.routingTables[method]
	pathname = strings.TrimSuffix(pathname, "/")
	existedHandler := routingTable.Search(pathname)
	if existedHandler != nil {
		panic("already exists")
	}
	routingTable.Insert(pathname, handler)
	return nil
}

func (r *Router) Get(pathname string, handler func(ctx *MyContext)) error {
	return r.register("get", pathname, handler)
}
func (r *Router) Post(pathname string, handler func(ctx *MyContext)) error {
	return r.register("post", pathname, handler)
}
func (r *Router) Patch(pathname string, handler func(ctx *MyContext)) error {
	return r.register("patch", pathname, handler)
}
func (r *Router) Put(pathname string, handler func(ctx *MyContext)) error {
	return r.register("put", pathname, handler)
}
func (r *Router) Delete(pathname string, handler func(ctx *MyContext)) error {
	return r.register("delete", pathname, handler)
}

func (h *Engine) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	fileServer := http.FileServer(http.Dir("./static"))

	ctx := NewMyContext(rw, r)

	ctx.Set("AuthUser", "test1111")

	routingTable := h.Router.routingTables[strings.ToLower(r.Method)]
	pathname := r.URL.Path
	pathname = strings.TrimSuffix(pathname, "/")

	fPath := path.Join("./static", pathname)
	fInfo, err := os.Stat(fPath)
	fExists := err == nil && !fInfo.IsDir()
	if fExists {
		fileServer.ServeHTTP(rw, r)
		return
	}

	targetNode := routingTable.Search(pathname)

	if targetNode == nil || targetNode.handler == nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	paramDicts := targetNode.ParseParams(r.URL.Path)

	ctx.SetParams(paramDicts)

	ch := make(chan struct{})
	panicCh := make(chan struct{})
	go func() {

		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				panicCh <- struct{}{}
			}
		}()
		time.Sleep(time.Second * 1)
		targetNode.handler(ctx)
		ch <- struct{}{}
	}()

	durationContext, cancel := context.WithTimeout(r.Context(), time.Second*3)
	defer cancel()
	select {
	case <-durationContext.Done():
		ctx.SetTimeout(true)
		fmt.Println("timeout")
		ctx.rw.Write([]byte("timeout"))
	case <-ch:
		fmt.Println("done")
	case <-panicCh:
		fmt.Println("panic")
		ctx.rw.WriteHeader(500)
	}

	return
}

func (e *Engine) Run() {
	http.ListenAndServe("localhost:8080", e)
}
