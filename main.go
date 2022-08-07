package main

import (
	"net/http"
	"tinygin/tinygin"
)

func main() {
	r := tinygin.NewEngine()

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

	r.Run()
}
