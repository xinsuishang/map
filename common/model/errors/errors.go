package errors

import "errors"

const (

	// System Code

	SuccessCode            = 0
	ParamErrCode           = 10001
	ServiceErrCode         = 10002
	ServiceNotFoundErrCode = 10003
	ServerHandleFailCode   = 10004

	// User ErrCode

)

type ErrNo struct {
	Code   int32
	ErrMsg string
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

func (e ErrNo) WithError(err error) ErrNo {
	e.ErrMsg = err.Error()
	return e
}

var (
	Success            = NewErrNo(SuccessCode, "Success")
	ParamErr           = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	ServiceErr         = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ServiceNotFoundErr = NewErrNo(ServiceNotFoundErrCode, "found service error")
	ServerHandleFail   = NewErrNo(ServerHandleFailCode, "ServerHandleFail")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
