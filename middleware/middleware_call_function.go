package _middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	_request "go-oceancode-core/model/request"
	_response "go-oceancode-core/model/response"
)

func MiddlewareCallFunction(r *ghttp.Request, reqPointer interface{}, handler func(ctx *_request.SessionContext) any) {
	sessionContext := &_request.SessionContext{
		UserId:    0,
		ProjectId: 0,
		TenantId:  0,
		Context:   r.Context(),
	}

	if reqPointer != nil {
		if err := r.Parse(reqPointer); err != nil {
			r.Response.WriteJson(_response.ResultOk(nil))
		} else {
			res := handler(sessionContext)
			r.Response.WriteJson(_response.ResultOk(&res))
		}
	} else {
		res := handler(sessionContext)
		r.Response.WriteJson(_response.ResultOk(res))
	}
}
