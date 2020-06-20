package main

import (
	".."
	"net/http"
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
		ctx.String(http.StatusOK, "hello %s", ctx.R.Host)
	})

	e.Get("/", func(ctx *tequila.Context) {
		ctx.String(http.StatusOK, "hello")
	})

	// 监听端口 等待处理handler
	_ = e.Run(":8099")
}
