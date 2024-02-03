package config

import "msp/common/model/config"

type ServerConfig struct {
	config.ServerConfig
	GatewayResource []GatewayConfig `mapstructure:"gateway_resource" json:"gateway_resource"`
}

type GatewayConfig struct {
	config.RpcSignConfig
	Route       string   `mapstructure:"route" json:"route"`
	FingerPrint string   `mapstructure:"finger_print" json:"finger_print"`
	ParentPath  string   `mapstructure:"parent_path" json:"parent_path"`
	IdlPath     string   `mapstructure:"idl_path" json:"idl_path"`
	IncludePath []string `mapstructure:"include_path" json:"include_path"`
}
