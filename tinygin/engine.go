package tinygin

import (
	"net/http"
	"strings"
)

type JS map[string]interface{}
type Engine struct {
	*Group
	groups []*Group
	router *Router
}

func NewEngine() *Engine {
	engine := &Engine{
		router: NewRouter(),
	}
	engine.Group = &Group{
		preStr:      "",
		parent:      nil,
		engine:      engine,
		middlewares: make([]HandlerFunc, 0),
	}
	engine.groups = []*Group{engine.Group}
	return engine
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.preStr) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	ctx := NewGinContext(w, req)
	ctx.handler = middlewares
	e.router.handle(ctx)
}
func (e *Engine) Run() error {
	return http.ListenAndServe(":8888", e)
}
