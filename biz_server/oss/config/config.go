package config

import "msp/common/model/config"

type ServerConfig struct {
	config.ServerConfig
	MysqlInfo config.MysqlConfig `mapstructure:"mysql" json:"mysql"`
}
