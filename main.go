package main

import (
	"go-web/action"
	"go-web/common"
	"go-web/middleware"
	server2 "go-web/server"
	"net/http"
)

func main() {

	// 创建服务
	server := server2.NewHttpServer("server", func(next middleware.Filter) middleware.Filter {
		return func(ctx *common.Context) {
		}
	})
	// 注册路由
	server.Route(http.MethodPost, "/signUp", action.SignUp)

	// 启动
	if err := server.Start(":8080"); err != nil {
		// 启动失败用panic停止服务
		panic(err)
	}
}
