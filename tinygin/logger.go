package tinygin

import (
	"log"
	"net/http"
	"time"
)

func Logger() HandlerFunc {
	return func(ctx *GinContext) {
		// Start timer
		t := time.Now()
		// Process request
		ctx.Next()
		// Calculate resolution time
		ctx.String(http.StatusOK, "\n[%d] %s in %v", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
		log.Printf("[%d] %s in %v", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}
