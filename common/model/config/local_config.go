package config

type LocalConfig struct {
	Name         string       `mapstructure:"name" json:"name"`
	LogConfig    LogConfig    `mapstructure:"log" json:"log"`
	ConsulConfig ConsulConfig `mapstructure:"consul" json:"consul"`
}
