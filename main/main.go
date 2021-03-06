package main

import (
	".."
	"log"
	"net/http"
	"time"
)

func main() {

	// 自定义Engine处理所有请求 engine实现的serveHTTP方法 等同于实现了handler的接口
	e := tequila.New()
	e.Post("/post", func(ctx *tequila.Context) {
		ctx.String(http.StatusOK, "this is  post method %s", ctx.R.URL)
	})

	// 返回一个json的新写法
	e.Get("/get", func(ctx *tequila.Context) {
		ctx.Json(http.StatusOK, tequila.J{
			"name": "joey",
			"age":  30,
		})
	})
	// 动态路由
	e.Get("/hello/:lang/say", func(ctx *tequila.Context) {
		ctx.String(http.StatusOK, "hello say %s", ctx.Param("lang"))
	})
	// 动态路由 *
	e.Get("/hello/*filePath", func(ctx *tequila.Context) {
		ctx.String(http.StatusOK, "hello give you file %s", ctx.Param("filePath"))
	})

	e.Get("/", func(ctx *tequila.Context) {
		ctx.String(http.StatusOK, "hello")
	})

	group := e.Group("/v1")
	group.Use(func(ctx *tequila.Context) {
		t := time.Now()
		ctx.Next()
		log.Printf("[%d] %s in %v", ctx.StatusCode, ctx.R.RequestURI, time.Since(t))
	})
	group.Get("/hello", func(ctx *tequila.Context) {
		ctx.String(http.StatusOK, "yoyoyoyo this is v1 %s", ctx.Path)
	})

	// 监听端口 等待处理handler
	_ = e.Run(":8099")
}
