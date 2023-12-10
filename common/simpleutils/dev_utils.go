package simpleutils

import (
	"msp/common/constant"
	"runtime"
)

func IsDev() bool {
	return runtime.GOOS != constant.Linux
}
