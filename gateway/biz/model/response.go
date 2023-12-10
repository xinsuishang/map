package model

import "msp/common/model/errors"

type Response struct {
	// Status code, 0-success, other values-failure
	Code int32 `form:"code" json:"code"`
	// Return status description
	Message string `form:"message" json:"message"`

	Data any `form:"data" json:"data,omitempty"`
}

func (r *Response) Success(data any) {
	r.Err(errors.Success)
	r.Data = data
}

func (r *Response) Err(err errors.ErrNo) {
	r.Code = err.Code
	r.Message = err.ErrMsg
}

func (r *Response) ParamsErrMsg(message string) {
	r.Code = errors.ParamErrCode
	r.Message = message
}

func (r *Response) BizErrMsg(message string) {
	r.Code = errors.ServiceErrCode
	r.Message = message
}
