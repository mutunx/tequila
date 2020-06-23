package tequila

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type J map[string]interface{}

type Context struct {
	W http.ResponseWriter
	R *http.Request
	// 请求数据
	Path   string
	Method string
	// 返回数据
	StatusCode int
	// 获取动态路由的参数
	Params map[string]string
	// 中间件
	handlers []handler
	index    int
}

// 创建
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W:      w,
		R:      r,
		Path:   r.URL.Path,
		Method: r.Method,
		index:  -1,
	}
}

func (c *Context) Next() {
	c.index++
	// 如果调用next就逐个获取handlers里面的handler处理 等处理完后再处理调用next方法之后的handler
	for ; c.index < len(c.handlers); c.index++ {
		c.handlers[c.index](c)
	}
}

// 获取参数
func (c *Context) Param(key string) string {
	return c.Params[key]
}

// 设置状态
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.W.WriteHeader(code)
}

// 设置请求头
func (c *Context) SetHeader(key, value string) {
	c.W.Header().Set(key, value)
}

// 获取参数 post get
func (c *Context) Query(key string) string {
	return c.R.URL.Query().Get(key)
}

func (c *Context) PostForm(key string) string {
	return c.R.PostForm.Get(key)
}

// 设置返回的格式
func (c *Context) Json(code int, j J) {
	// 设置返回头
	c.W.Header().Set("Content-Type", "application/json; charset=utf-8")
	// 设置返回状态码
	c.W.WriteHeader(code)
	// 设置返回数据
	encode := json.NewEncoder(c.W)
	if err := encode.Encode(j); err != nil {
		http.NotFound(c.W, c.R)
	}
}

func (c *Context) String(code int, format string, value ...interface{}) {
	c.W.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.W.WriteHeader(code)
	_, _ = c.W.Write([]byte(fmt.Sprintf(format, value...)))
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	_, _ = c.W.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html; charset=utf-8")
	c.Status(code)
	_, _ = c.W.Write([]byte(html))
}
