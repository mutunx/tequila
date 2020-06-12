package main

import (
	"fmt"
	"net/http"
)

type Engine struct {
}

func main() {

	// 自定义Engine处理所有请求 engine实现的serveHTTP方法 等同于实现了handler的接口
	e := new(Engine)

	// 监听端口 等待处理handler
	_ = http.ListenAndServe(":8099", e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 根据请求不同的请求地址进行不同的处理
	switch r.URL.Path {
	case "/":
		_, _ = fmt.Fprintf(w, "%s %s", r.Method, r.URL)
	case "/hello":
		_, _ = fmt.Fprintf(w, "%s %s hello!", r.Method, r.URL)
	}
}
