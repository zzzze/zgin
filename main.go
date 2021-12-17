package main

import (
	"net/http"
	"zgin"
)

func main() {
	r := zgin.New()
	r.GET("/", func(c *zgin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})
	r.GET("/hello", func(c *zgin.Context) {
		// expect /hello?name=zhangsan
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.GET("/hello/:name", func(c *zgin.Context) {
		// expect /hello/zhangsan
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Params["name"], c.Path)
	})
	r.POST("/login", func(c *zgin.Context) {
		c.JSON(http.StatusOK, zgin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
