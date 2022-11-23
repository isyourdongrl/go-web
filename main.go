package main

import (
	"net/http"
)

func main() {

	// 创建服务
	server := NewHttpServer("server")
	// 注册路由
	server.Route(http.MethodPost, "/signUp", SignUp)

	// 启动
	if err := server.Start(":8080"); err != nil {
		// 启动失败用panic停止服务
		panic(err)
	}
}
