package initialize

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
	"msp/gateway/config"
	"net"
	"strconv"
)

// InitRegistry to init consul
func InitRegistry(Port int) (registry.Registry, *registry.Info) {

	// build a consul client
	cfg := api.DefaultConfig()
	cfg.Address = net.JoinHostPort(
		config.GlobalLocalConfig.ConsulConfig.Host,
		strconv.Itoa(config.GlobalLocalConfig.ConsulConfig.Port))
	consulClient, err := api.NewClient(cfg)
	if err != nil {
		hlog.Fatalf("new consul client failed: %s", err.Error())
	}

	r := consul.NewConsulRegister(consulClient,
		consul.WithCheck(&api.AgentServiceCheck{
			Interval:                       "7s",
			Timeout:                        "4s",
			DeregisterCriticalServiceAfter: "15s",
		}))

	// Using snowflake to generate service name.
	sf, err := snowflake.NewNode(2)
	if err != nil {
		hlog.Fatalf("generate service name failed: %s", err.Error())
	}
	info := &registry.Info{
		ServiceName: config.GlobalServerConfig.Name,
		Addr: utils.NewNetAddr("tcp", net.JoinHostPort(config.GlobalServerConfig.Host,
			strconv.Itoa(Port))),
		Tags: map[string]string{
			"ID": sf.Generate().Base36(),
		},
		Weight: registry.DefaultWeight,
	}
	return r, info
}
