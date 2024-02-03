package errors

import (
	"github.com/cloudwego/kitex/pkg/kerrors"
)

const (

	// System Code
	SuccessCode            = 0
	ServerHandleFailCode   = 500
	ParamErrCode           = 10001
	ServiceErrCode         = 10002
	ServiceNotFoundErrCode = 10003

	// User ErrCode

)

func NewErrNo(code int32, msg string) kerrors.BizStatusErrorIface {
	return kerrors.NewBizStatusError(code, msg)
}

var (
	Success            = NewErrNo(SuccessCode, "Success")
	ParamErr           = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	ServiceErr         = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ServiceNotFoundErr = NewErrNo(ServiceNotFoundErrCode, "found service error")
	ServerHandleFail   = NewErrNo(ServerHandleFailCode, "ServerHandleFail")
)
