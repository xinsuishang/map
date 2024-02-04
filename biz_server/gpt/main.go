package main

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"msp/biz_server/gpt/config"
	"msp/biz_server/gpt/initialize"
	"msp/biz_server/gpt/internal/infra/mysql/model"
	"msp/biz_server/kitex_gen/gpt"
	"msp/biz_server/kitex_gen/gpt/chatservice"
	"msp/common/constant"
	commoninit "msp/common/initialize"
	"msp/common/valid"
	"net"
)

func main() {
	// 本地配置文件初始化
	initialize.InitLocalConfig()
	// 根据本地配置文件初始化日志
	commoninit.InitLogger(&config.GlobalLocalConfig.LogConfig, config.GlobalLocalConfig.Name, config.IsDev)
	// 根据本地配置文件初始化配置
	initialize.InitConfig()
	// 根据本地配置文件初始化数据库
	db := initialize.InitDB()
	// 注册中心
	r, info := initialize.InitRegistry(config.GlobalServerConfig.Port, config.GlobalServerConfig.Group, config.GlobalServerConfig.Version)
	//p := provider.NewOpenTelemetryProvider(
	//	provider.WithServiceName(config.GlobalServerConfig.Name),
	//	//provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
	//	provider.WithInsecure(),
	//)
	//defer p.Shutdown(context.Background())

	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", config.GlobalServerConfig.Port))
	svr := chatservice.NewServer(func() gpt.ChatService {
		return &ChatServiceImpl{
			repository: model.NewEntClient(db),
		}
	}(),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithSuite(valid.NewServerSuite()),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name,
			Tags: map[string]string{
				constant.Group:   "production",
				constant.Version: "2.0",
			}}))

	err := svr.Run()

	if err != nil {
		klog.Fatal(err.Error())
	}
}
