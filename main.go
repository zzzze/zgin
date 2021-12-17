package main

import (
	"log"
	"net/http"
	"time"
	"zgin"
)

func onlyForV2() zgin.HandlerFunc {
	return func(c *zgin.Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := zgin.New()
	r.Use(zgin.Logger())
	r.GET("/index", func(c *zgin.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *zgin.Context) {
			c.HTML(http.StatusOK, "<h1>Hello World</h1>")
		})
		v1.GET("/hello", func(c *zgin.Context) {
			// expect /hello?name=zhangsan
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *zgin.Context) {
			// expect /hello/zhangsan
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Params["name"], c.Path)
		})
		v2.POST("/login", func(c *zgin.Context) {
			println("ggg")
			c.JSON(http.StatusOK, zgin.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}
	r.Run(":9999")
}
