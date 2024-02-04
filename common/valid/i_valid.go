package valid

import "msp/common/model/errors"

type ValidIface interface {
	IsValid() error
}

// Validate 校验参数并包装成bizError
func Validate(i ValidIface) error {
	err := i.IsValid()
	if err != nil {
		return errors.NewErrNo(errors.ParamErrCode, err.Error())
	}
	return nil
}
