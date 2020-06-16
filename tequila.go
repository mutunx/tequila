package tequila

import (
	"net/http"
)

type handler func(ctx *Context)

// 同处理器统一处理请求
type Engine struct {
	router *router
}

// 创建新的路由
func New() *Engine {
	return &Engine{router: newRouter()}
}

// 开发 post 和 get 请求
func (e *Engine) Get(path string, h handler) {
	e.addRoute("GET", path, h)
}

func (e *Engine) Post(path string, h handler) {
	e.addRoute("POST", path, h)
}

// 添加路由  主程序的添加路由由主程序统一处理 发给router
func (e *Engine) addRoute(method, path string, h handler) {
	e.router.addRoute(method, path, h)
}

// 处理请求 统一交给router处理
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	cxt := newContext(w, r)
	e.router.handle(cxt)
}

// 监听端口
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
