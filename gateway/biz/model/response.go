package model

import (
	"github.com/cloudwego/kitex/pkg/kerrors"
	"msp/common/model/errors"
)

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

func (r *Response) Err(err error) {
	bizErr, ok := kerrors.FromBizStatusError(err)
	if !ok {
		r.Err(errors.NewErrNo(errors.ServerHandleFailCode, err.Error()))
		return
	}
	r.Code = bizErr.BizStatusCode()
	r.Message = bizErr.BizMessage()
}

func (r *Response) ParamsErrMsg(message string) {
	r.Err(errors.NewErrNo(errors.ParamErrCode, message))
}

func (r *Response) BizErrMsg(message string) {
	r.Err(errors.NewErrNo(errors.ServiceErrCode, message))
}
