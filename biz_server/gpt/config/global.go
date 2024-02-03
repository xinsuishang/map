package config

import "msp/common/model/config"

var (
	IsDev bool

	GlobalServerConfig ServerConfig // 持有通过consul获取的全部服务配置，包括本服务的ip、端口、providerName

	GlobalLocalConfig config.LocalConfig
)
