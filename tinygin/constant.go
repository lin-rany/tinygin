package tinygin

import (
	"net/http"
)

func PathNotFound(ctx *GinContext) {
	ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
}
