package gfJson

import (
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (res *Response) JSON(r *ghttp.Request) error {
	return  r.Response.WriteJsonExit(res)
}

// 操作成功
func (res *Response) Success(r *ghttp.Request)  {
	res.Code = http.StatusOK
	res.Msg = "操作成功！"
	res.Data = nil

	_ = res.JSON(r)
}

// 数据错误
func (res *Response) ServerError(r *ghttp.Request, data interface{})  {
	res.Code = http.StatusInternalServerError
	res.Msg = "数据错误！"
	res.Data = data

	_ = res.JSON(r)
}

// 未授权
func (res *Response) UnAuthorized(r *ghttp.Request)  {
	res.Code = http.StatusUnauthorized
	res.Msg = "未授权！"

	_ = res.JSON(r)
}

// 错误的请求
func (res *Response) BadRequest(r *ghttp.Request)  {
	res.Code = http.StatusBadRequest
	res.Msg = "错误的请求！"

	_ = res.JSON(r)
}

