package tequila

import (
	"fmt"
	"net/http"
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

	cxt := newContext(w, r)
	e.router.handle(cxt)
}

// 监听端口
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
