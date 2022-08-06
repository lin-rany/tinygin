package tinygin

import (
	"fmt"
	"log"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func NewEngine() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}
func (e *Engine) AddRouter(method, path string, handler HandlerFunc) {
	log.Printf("AddRouter method %v path %v handler %v", method, path, handler)
}
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
func (e *Engine) Run() error {
	return http.ListenAndServe(":8888", e)
}
