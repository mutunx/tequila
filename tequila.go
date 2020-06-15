package tequila

import (
	"fmt"
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request)

// 同处理器统一处理请求
type Engine struct {
	router map[string]handler
}

// 创建新的路由
func New() *Engine {
	return &Engine{
		router: make(map[string]handler),
	}
}

// 添加路由
func (e *Engine) addRoute(method, path string, h handler) {
	// 将方法与地址绑定 用于区分不同方法相同路径的handler
	key := fmt.Sprintf("%s-%s", method, path)
	e.router[key] = h
}

// 开发 post 和 get 请求
func (e *Engine) Get(path string, h handler) {
	e.addRoute("GET", path, h)
}

func (e *Engine) Post(path string, h handler) {
	e.addRoute("POST", path, h)
}

// 处理请求
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 获取方法和地址 用于拼装key
	key := fmt.Sprintf("%s-%s", r.Method, r.URL.Path)
	// 查看是否在路由中
	h, ok := e.router[key]
	if ok {
		h(w, r)
	} else {
		_, _ = fmt.Fprintf(w, "404 not found %s", r.URL.Path)
	}
}

// 监听端口
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
