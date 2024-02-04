package initialize

import (
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"msp/biz_server/gpt/config"
	"msp/common/simpleutils"
	"net"
	"strconv"
)

func InitLocalConfig() {
	v := viper.New()
	v.SetConfigFile("./biz_server/gpt/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		hlog.Fatalf("read viper config failed: %s", err.Error())
	}
	if err := v.Unmarshal(&config.GlobalLocalConfig); err != nil {
		hlog.Fatalf("unmarshal err failed: %s", err.Error())
	}
	hlog.Infof("Config Info: %v", config.GlobalLocalConfig)

	config.IsDev = simpleutils.IsDev()
}

func InitConfig() {
	consulConfig := config.GlobalLocalConfig.ConsulConfig
	cfg := api.DefaultConfig()
	cfg.Address = net.JoinHostPort(
		consulConfig.Host,
		strconv.Itoa(consulConfig.Port))
	consulClient, err := api.NewClient(cfg)
	if err != nil {
		hlog.Fatalf("new consul client failed: %s", err.Error())
	}
	content, _, err := consulClient.KV().Get(consulConfig.Key, nil)
	if err != nil {
		hlog.Fatalf("consul kv failed: %s", err.Error())
	}

	err = sonic.Unmarshal(content.Value, &config.GlobalServerConfig)
	if err != nil {
		hlog.Fatalf("sonic unmarshal config failed: %s", err.Error())
	}
}
