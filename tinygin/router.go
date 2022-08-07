package tinygin

import (
	"log"
	"strings"
)

type HandlerFunc func(ctx *GinContext)
type Router struct {
	root     map[string]*node
	handlers map[string]HandlerFunc
}

func GetHandlerKey(method, path string) string {
	return method + "_" + path
}
func NewRouter() *Router {
	return &Router{
		root:     make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *Router) handle(ctx *GinContext) {
	knode, parms := r.GetRoute(ctx.Method, ctx.Path)
	if knode != nil {
		ctx.Params = parms
		r.handlers[GetHandlerKey(ctx.Method, knode.path)](ctx)
	} else {
		PathNotFound(ctx)
	}
}
func (r *Router) AddRoute(method, path string, handler HandlerFunc) {
	log.Printf("AddRoute method %v path %v handler %v", method, path, handler)
	_, ok := r.root[method]
	if !ok {
		r.root[method] = &node{}
	}
	r.root[method].Insert(path, Parseparts(path), 0)
	r.handlers[GetHandlerKey(method, path)] = handler
}
func (r *Router) GetRoute(method, path string) (*node, map[string]string) {
	if _, ok := r.root[method]; !ok {
		return nil, nil
	}
	knode := r.root[method].Search(path, Parseparts(path), 0)
	if knode == nil {
		return nil, nil
	}
	params := make(map[string]string)
	parts := Parseparts(path)
	knodeparts := Parseparts(knode.path)
	for index, part := range knodeparts {
		if part[0] == ':' {
			params[part[1:]] = parts[index]
		}
		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(parts[index:], "/")
			break
		}
	}
	return knode, params
}
