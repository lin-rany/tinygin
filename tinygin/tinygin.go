package tinygin

import (
	"net/http"
)

type JS map[string]interface{}
type Engine struct {
	router *Router
}

func NewEngine() *Engine {
	return &Engine{
		router: NewRouter(),
	}
}
func (e *Engine) GET(path string, handler HandlerFunc) {
	e.router.AddRoute("GET", path, handler)
}
func (e *Engine) POST(path string, handler HandlerFunc) {
	e.router.AddRoute("POST", path, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	cxt := NewGinContext(w, req)
	e.router.handle(cxt)
}
func (e *Engine) Run() error {
	return http.ListenAndServe(":8888", e)
}
