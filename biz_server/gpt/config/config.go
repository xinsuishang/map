package config

import "msp/common/model/config"

type ServerConfig struct {
	config.RpcServerConfig
	MysqlInfo config.MysqlConfig `mapstructure:"mysql" json:"mysql"`
}
