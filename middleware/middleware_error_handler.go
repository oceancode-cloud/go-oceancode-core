package _middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	_response "go-oceancode-core/model/response"
)

func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		r.Response.ClearBuffer()
		r.Response.WriteJson(_response.ResultError())
	}
}
