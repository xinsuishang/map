package config

type LogConfig struct {
	LogDir   string `mapstructure:"dir" json:"dir"`
	LogDebug bool   `mapstructure:"level" json:"level"`
}
