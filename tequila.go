package tequila

import (
	"fmt"
	"net/http"
	"strings"
)

type handler func(ctx *Context)

// 分组处理
type RouterGroup struct {
	prefix     string    // 分组的标识符
	middleware []handler // 中间件
	engine     *Engine   // 所有分组共享一个实例
}

// 同处理器统一处理请求
type Engine struct {
	*RouterGroup // 继承分组的所有方法
	router       *router
	groups       []*RouterGroup //所有分组
}

// 创建新的路由
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{
		engine: engine,
	}
	// 添加到分组中
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// 添加方法到中间件中
func (group *RouterGroup) Use(h handler) {
	group.middleware = append(group.middleware, h)
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	// 获取engine
	engine := group.engine
	routerGroup := &RouterGroup{
		prefix:     prefix,
		middleware: nil,
		engine:     engine,
	}
	// 添加新的分组到系统分组中
	engine.groups = append(engine.groups, routerGroup)
	return routerGroup
}

// 开发 post 和 get 请求
func (group *RouterGroup) Get(path string, h handler) {
	group.addRoute("GET", path, h)
}

func (group *RouterGroup) Post(path string, h handler) {
	group.addRoute("POST", path, h)
}

// 添加路由  主程序的添加路由由主程序统一处理 发给router
func (group *RouterGroup) addRoute(method, path string, h handler) {
	// 增加分组的信息
	path = group.prefix + path
	fmt.Printf("router %s -%s\n", method, path)
	group.engine.router.addRoute(method, path, h)
}

// 处理请求 统一交给router处理
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 获取当前分组 获取当前分组中间件 交给context处理
	var middlewares []handler
	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middleware...)
		}
	}
	ctx := newContext(w, r)
	ctx.handlers = middlewares
	e.router.handle(ctx)
}

// 监听端口
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
