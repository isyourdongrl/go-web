package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

// ReadJson 将 http json请求参数反序列化
func (c *Context) ReadJson(req interface{}) error {
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, req); err != nil {
		return err
	}
	return nil
}

// WriteJson 统一返回
func (c *Context) WriteJson(code int, resp interface{}) (err error) {
	// 写入状态码
	c.W.WriteHeader(code)

	// 序列化响应
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	// 将响应写入w
	if _, err = c.W.Write(respJson); err != nil {
		return err
	}
	return err
}

func (c *Context) OkHttp(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}

func (c *Context) SystemErrJson(resp interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, resp)
}

func (c *Context) BadRequestJson(resp interface{}) error {
	return c.WriteJson(http.StatusBadRequest, resp)
}
