package initialize

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/hashicorp/consul/api"
	consul "github.com/kitex-contrib/registry-consul"
	"msp/biz_server/oss/config"
	"msp/common/constant"
	"net"
	"strconv"
)

// InitRegistry to init consul
func InitRegistry(port int, group, version string) (registry.Registry, *registry.Info) {
	r, err := consul.NewConsulRegister(net.JoinHostPort(
		config.GlobalLocalConfig.ConsulConfig.Host,
		strconv.Itoa(config.GlobalLocalConfig.ConsulConfig.Port)),
		consul.WithCheck(&api.AgentServiceCheck{
			Interval:                       "7s",
			Timeout:                        "4s",
			DeregisterCriticalServiceAfter: "15s",
		}))
	if err != nil {
		klog.Fatalf("new consul register failed: %s", err.Error())
	}

	// Using snowflake to generate service name.
	sf, err := snowflake.NewNode(2)
	if err != nil {
		klog.Fatalf("generate service name failed: %s", err.Error())
	}
	info := &registry.Info{
		ServiceName: config.GlobalServerConfig.Name,
		Addr:        utils.NewNetAddr("tcp", net.JoinHostPort(config.GlobalServerConfig.Host, strconv.Itoa(port))),
		Tags: map[string]string{
			"ID":             sf.Generate().Base36(),
			constant.Group:   group,
			constant.Version: version,
		},
	}
	return r, info
}
