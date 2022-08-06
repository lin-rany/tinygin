package tinygin

import "net/http"

type GinContext struct {
	Writer http.ResponseWriter
	Req    *http.Request
	// writer

	// req
}
