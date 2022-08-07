package tinygin

import (
	"log"
	"net/http"
)

type HandlerFunc func(ctx *GinContext)
type Router struct {
	handlers map[string]HandlerFunc
}

func GetHandlerKey(method, path string) string {
	return method + "_" + path
}
func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
	}
}
func (r *Router) handle(ctx *GinContext) {
	if handler, ok := r.handlers[GetHandlerKey(ctx.Method, ctx.Path)]; ok {
		handler(ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
	}
}
func (r *Router) AddRouter(method, path string, handler HandlerFunc) {
	log.Printf("AddRouter method %v path %v handler %v", method, path, handler)
	r.handlers[GetHandlerKey(method, path)] = handler

}
