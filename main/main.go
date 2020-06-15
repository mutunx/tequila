package main

import (
	".."
	"fmt"
	"net/http"
)

type Engine struct {
}

func main() {

	// 自定义Engine处理所有请求 engine实现的serveHTTP方法 等同于实现了handler的接口
	e := tequila.New()
	e.Post("/testPost", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "this is  post method %s", r.URL)
	})

	e.Get("/testGet", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "this is get method %s", r.URL.Path)
	})

	e.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "hello")
	})

	// 监听端口 等待处理handler
	_ = e.Run(":8099")
}
