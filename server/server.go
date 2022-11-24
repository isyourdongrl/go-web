package server

import (
	"go-web/common"
	"go-web/middleware"
	"net/http"
)

// Server 将服务抽象化，每个服务都有路由和启动
type Server interface {
	middleware.Routable
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler middleware.Handler
	root    middleware.Filter
}

// Route 向外暴露出注册路由的方法，只需要传参数为context的函数即可
func (this *sdkHttpServer) Route(method string, pattern string, handleFunc middleware.HandlerFunc) {
	this.handler.Route(method, pattern, handleFunc)
}

// Start 开启服务
func (this *sdkHttpServer) Start(address string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c := common.NewContext(writer, request)
		this.root(c)
	})
	http.Handle("/", this.handler)
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string, builders ...middleware.FilterBuilder) Server {
	// 创建handler
	handler := middleware.NewHandlerBasedOnMap()
	// 定义filter的根结点
	var root middleware.Filter = func(ctx *common.Context) {
		handler.ServeHTTP(ctx.W, ctx.R)
	}

	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}

	return &sdkHttpServer{
		Name:    name,
		handler: handler,
		root:    root,
	}
}
