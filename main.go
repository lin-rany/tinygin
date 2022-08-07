package main

import (
	"net/http"
	"tinygin/tinygin"
)

func main() {
	svr := tinygin.NewEngine()
	svr.Get("/hello", func(ctx *tinygin.GinContext) {
		ctx.String(http.StatusOK, "hello world")
	})
	svr.Get("/login", func(ctx *tinygin.GinContext) {
		ctx.JSON(http.StatusOK, tinygin.JS{
			"username": "ykw",
			"password": 1234,
		})
	})

	svr.Run()
}
