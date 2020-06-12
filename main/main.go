package main

import (
	"fmt"
	"net/http"
)

func main() {

	// handleFunc 会把方法转成赋给一个handleFunc函数变量 这个变量实现了serveHTTP在里面调用本方法
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, "%s", request.Method)
	})

	// 监听端口 等待处理handler
	_ = http.ListenAndServe(":8099", nil)
}
