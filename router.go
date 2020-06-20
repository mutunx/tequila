package tequila

import (
	"fmt"
	"net/http"
	"strings"
)

type router struct {
	handlers map[string]handler
	roots    map[string]*node // 区分不同请求方法的前缀树
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]handler),
		roots:    make(map[string]*node),
	}
}

// 添加路由
func (r *router) addRoute(method, path string, h handler) {
	// 将地址中的节点添加到树node中
	parts := parsePath(path)
	child := r.roots[method]
	if child == nil {
		child = &node{}
		r.roots[method] = child
	}
	child.insert(path, parts, 0)
	// 将方法与地址绑定 用于区分不同方法相同路径的handler
	key := fmt.Sprintf("%s-%s", method, path)
	r.handlers[key] = h
}

func parsePath(path string) []string {
	return strings.Split(path, "/")
}

// 统一处理方法
func (r *router) handle(ctx *Context) {
	n := r.getRoute(ctx.Method, ctx.Path)
	if n != nil {
		// 获取方法和地址 用于拼装key  根据最后匹配的地址进行操作 用于匹配动态路由
		key := fmt.Sprintf("%s-%s", ctx.Method, n.Path)
		r.handlers[key](ctx)

	} else {
		http.NotFound(ctx.W, ctx.R)
	}
}

/**
输入:请求的方法和请求地址
返回:节点||nil
用于查找所请求的方法和请求的地址是否存在,存在则返回节点,不存在则返回nil
*/
func (r *router) getRoute(method, path string) *node {
	// 将地址中的节点添加到树node中
	parts := parsePath(path)
	child, ok := r.roots[method]
	if !ok {
		return nil
	}
	return child.search(parts, 0)
}
