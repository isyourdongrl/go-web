package main

import (
	"net/http"
)

// Server 将服务抽象化，每个服务都有路由和启动
type Server interface {
	Routable
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler Handler
	root    Filter
}

// SignUpReq 注册请求参数
type SignUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

// CommonResponse 统一返回参数
type CommonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// Route 向外暴露出注册路由的方法，只需要传参数为context的函数即可
func (this *sdkHttpServer) Route(method string, pattern string, handleFunc HandlerFunc) {
	this.handler.Route(method, pattern, handleFunc)
}

// Start 开启服务
func (this *sdkHttpServer) Start(address string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c := NewContext(writer, request)
		this.root(c)
	})
	http.Handle("/", this.handler)
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string, builders ...FilterBuilder) Server {
	handler := NewHandlerBasedOnMap()
	// 定义filter的根结点
	var root Filter = func(ctx *Context) {
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

func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		R: request,
		W: writer,
	}
}

// SignUp 注册
func SignUp(ctx *Context) {
	// 注册参数
	req := &SignUpReq{}

	if err := ctx.ReadJson(req); err != nil {
		ctx.BadRequestJson(err)
		return
	}

	resp := &CommonResponse{
		Msg:  "success",
		Data: 123,
	}
	if err := ctx.WriteJson(200, resp); err != nil {
		ctx.SystemErrJson(err)
		return
	}
	return
}
