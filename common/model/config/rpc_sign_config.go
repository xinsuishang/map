package config

type RpcSignConfig struct {
	SvrName string `mapstructure:"svr_name" json:"svr_name"`
	Group   string `mapstructure:"group" json:"group"`
	Version string `mapstructure:"version" json:"version"`
}
