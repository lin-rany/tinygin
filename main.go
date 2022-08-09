package main

import (
	"math/rand"
	"net/http"
	"time"
	"tinygin/tinygin"
)

func main() {
	r := tinygin.NewEngine()
	r.Use(tinygin.Logger())
	r.GET("/hello", func(c *tinygin.GinContext) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *tinygin.GinContext) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.GetParam("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *tinygin.GinContext) {
		c.JSON(http.StatusOK, tinygin.JS{"filepath": c.GetParam("filepath")})
	})
	v1 := r.NewGroup("/v1")
	v1.GET("/where", func(ctx *tinygin.GinContext) {
		ctx.String(http.StatusOK, "you path is %s", ctx.Path)
	})

	v2 := v1.NewGroup("/v2")
	v2.GET("/where", func(ctx *tinygin.GinContext) {

		time.Sleep(time.Duration(rand.Int31n(10000)) * time.Millisecond)
		ctx.String(http.StatusOK, "you path is %s", ctx.Path)
	})

	r.Run()
}
