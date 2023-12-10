package simpleutils

import (
	"fmt"
	"msp/common/constant"
)

func MySqlDSNFormatUtil(user, password, host string, port int, dbName string) string {
	return fmt.Sprintf(constant.MySqlDSNFormat, user, password, host, port, dbName)
}
