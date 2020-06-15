package main

import (
	".."
	"encoding/json"
	"fmt"
	"net/http"
)

type Engine struct {
}

func main() {

	// 自定义Engine处理所有请求 engine实现的serveHTTP方法 等同于实现了handler的接口
	e := tequila.New()
	e.Post("/post", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "this is  post method %s", r.URL)
	})

	// 返回一个json的原生写法
	e.Get("/get", func(w http.ResponseWriter, r *http.Request) {
		// 生成json对象
		obj := map[string]interface{}{
			"name": "joey",
			"age":  30,
		}
		// 设置返回格式 返回状态
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(obj); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	e.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "hello")
	})

	// 监听端口 等待处理handler
	_ = e.Run(":8099")
}
