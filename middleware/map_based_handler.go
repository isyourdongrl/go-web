package middleware

import (
	"go-web/common"
	"net/http"
)

// Routable Server和Handler都有注册路由的动作，提出来
type Routable interface {
	Route(method string, pattern string, handleFunc HandlerFunc)
}

type Handler interface {
	http.Handler
	Routable
}

type HandlerBasedOnMap struct {
	// 控制器，key是 method+url ，val是路由函数
	handlers map[string]func(ctx *common.Context)
}

// ServeHTTP 执行路由
func (h *HandlerBasedOnMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := h.key(r.Method, r.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		// 有这个路由就执行
		handler(common.NewContext(w, r))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found this Page ~"))
	}
}

// Route 注册路由
// key：method + "#" + pattern
// value：func(ctx *Context)
func (h *HandlerBasedOnMap) Route(method string, pattern string, handleFunc HandlerFunc) {
	key := h.key(method, pattern)
	h.handlers[key] = handleFunc
}

func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	return method + "#" + pattern
}

// 判断 HandlerBasedOnMap 是否实现了 Handler 接口
var _ Handler = &HandlerBasedOnMap{}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		handlers: make(map[string]func(ctx *common.Context)),
	}
}
