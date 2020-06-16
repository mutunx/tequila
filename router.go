package tequila

import (
	"fmt"
	"net/http"
)

type router struct {
	handlers map[string]handler
}

func newRouter() *router {
	return &router{handlers: make(map[string]handler)}
}

// 添加路由
func (r *router) addRoute(method, path string, h handler) {
	// 将方法与地址绑定 用于区分不同方法相同路径的handler
	key := fmt.Sprintf("%s-%s", method, path)
	r.handlers[key] = h
}

// 统一处理方法
func (r *router) handle(ctx *Context) {
	// 获取方法和地址 用于拼装key
	key := fmt.Sprintf("%s-%s", ctx.Method, ctx.Path)
	if handler := r.handlers[key]; handler != nil {
		handler(ctx)
	} else {
		http.NotFound(ctx.W, ctx.R)
	}
}
