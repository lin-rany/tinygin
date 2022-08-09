package tinygin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GinContext struct {
	// origin
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
	// middleware
	index   int
	handler []HandlerFunc
}

func NewGinContext(writer http.ResponseWriter, req *http.Request) *GinContext {
	return &GinContext{
		Writer: writer,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}
func (c *GinContext) Next() {
	c.index++
	hlen := len(c.handler)
	for ; c.index < hlen; c.index++ {
		c.handler[c.index](c)
	}
}

func (c *GinContext) GetParam(key string) string {
	if val, ok := c.Params[key]; ok {
		return val
	} else {
		return ""
	}
}
func (c *GinContext) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *GinContext) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *GinContext) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *GinContext) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *GinContext) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}
