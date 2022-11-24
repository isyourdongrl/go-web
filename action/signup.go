package action

import (
	"go-web/common"
)

// SignUpReq 注册请求参数
type SignUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

// SignUp 注册
func SignUp(ctx *common.Context) {
	// 注册参数
	req := &SignUpReq{}

	if err := ctx.ReadJson(req); err != nil {
		ctx.BadRequestJson(err)
		return
	}

	resp := &common.CommonResponse{
		Msg:  "success",
		Data: 123,
	}
	if err := ctx.WriteJson(200, resp); err != nil {
		ctx.SystemErrJson(err)
		return
	}
	return
}
