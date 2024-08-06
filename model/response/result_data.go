package _response

import _const "go-oceancode-core/model/const"

type ResultData struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}

func ResultOk(data interface{}) (res *ResultData) {
	return &ResultData{
		Code:    _const.SUCCESS,
		Message: "SUCCESS",
		Results: data,
	}
}

func ResultError() (res *ResultData) {
	return &ResultData{
		Code:    _const.ERROR,
		Message: "ERROR",
		Results: nil,
	}
}

func (r *ResultData) SetData(data interface{}) (res *ResultData) {
	r.Results = data
	return r
}

func (r *ResultData) SetMessage(message string) (res *ResultData) {
	r.Message = message
	return r
}

func (r *ResultData) SetCode(code string) (res *ResultData) {
	r.Code = code
	return r
}

func (r *ResultData) GetData() interface{} {
	return r.Results
}

func (r *ResultData) GetCode() string {
	return r.Code
}

func (r *ResultData) IsSuccess() bool {
	return r.Code == _const.SUCCESS
}
