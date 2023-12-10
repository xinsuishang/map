package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"msp/biz_server/kitex_gen/oss"
	"msp/biz_server/kitex_gen/oss/uploadservice"
	"msp/biz_server/oss/config"
	"msp/biz_server/oss/initialize"
	"msp/biz_server/oss/internal/infra/mysql/model"
)

func main() {
	// 本地配置文件初始化
	initialize.InitLocalConfig()
	// 根据本地配置文件初始化日志
	initialize.InitLogger()
	// 根据本地配置文件初始化配置
	initialize.InitConfig()
	// 根据本地配置文件初始化数据库
	db := initialize.InitDB()
	// 注册中心
	r, info := initialize.InitRegistry(config.GlobalServerConfig.Port)
	//p := provider.NewOpenTelemetryProvider(
	//	provider.WithServiceName(config.GlobalServerConfig.Name),
	//	//provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
	//	provider.WithInsecure(),
	//)
	//defer p.Shutdown(context.Background())

	model.NewEntClient(db)

	svr := uploadservice.NewServer(func() oss.UploadService {
		return &UploadServiceImpl{
			repository: model.NewEntClient(db),
		}
	}(),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}))

	err := svr.Run()

	if err != nil {
		klog.Fatal(err.Error())
	}
}
