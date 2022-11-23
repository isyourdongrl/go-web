package main

import "net/http"

type Routable interface {
	Route(method string, pattern string, handleFunc HandlerFunc)
}

type Handler interface {
	http.Handler
	Routable
}

type HandlerBasedOnMap struct {
	// 控制器，key是method+url ，val是路由函数
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBasedOnMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := h.key(r.Method, r.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		// 有这个路由就执行
		handler(NewContext(w, r))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found this Page ~"))
	}
}

// Route 注册路由
func (h *HandlerBasedOnMap) Route(method string, pattern string, handleFunc HandlerFunc) {
	key := h.key(method, pattern)
	h.handlers[key] = handleFunc
}

func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	return method + "#" + pattern
}

var _ Handler = &HandlerBasedOnMap{}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		handlers: make(map[string]func(ctx *Context)),
	}
}
