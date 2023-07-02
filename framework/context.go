package framework

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MyContext struct {
	rw http.ResponseWriter
	r  *http.Request
}

func NewMyContext(rw http.ResponseWriter, r *http.Request) *MyContext {
	return &MyContext{
		rw: rw,
		r:  r,
	}
}

func (ctx *MyContext) Json(data any) {
	responseData, err := json.Marshal(data)
	if err != nil {
		ctx.rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.rw.Header().Set("Content-Type", "application-json")
	ctx.rw.WriteHeader(http.StatusOK)
	ctx.rw.Write(responseData)
}

func (ctx *MyContext) WriteString(data string) {
	ctx.rw.WriteHeader(http.StatusOK)
	fmt.Fprint(ctx.rw, data)
}

func (ctx *MyContext) QueryAll() map[string][]string {
	return ctx.r.URL.Query()
}

func (ctx *MyContext) QueryKey(key string, defaultValue string) string {
	values := ctx.QueryAll()

	if target, ok := values[key]; ok {
		if len(target) == 0 {
			return defaultValue
		}

		return target[len(target)-1]
	}

	return defaultValue
}
