## make some progress

```shell
go install github.com/cloudwego/hertz/cmd/hz@latest
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
go install github.com/cloudwego/thriftgo@latest
go install github.com/cloudwego/thrift-gen-validator@latest
go install github.com/oligot/go-mod-upgrade@latest
```

```shell
docker run \
-d \
-p 8300:8300 \
-p 8301:8301 \
-p 8302:8302 \
-p 8500:8500 \
-p 8501:8501 \
-p 8600:8600/udp \
--name=consule \
consul:1.15.2 agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0

createa msp go project
cd msp
mkdir gateway
mkdir biz_server
mkdir idl

cd gateway
hz new

cd ../biz_server
kitex --thrift-plugin validator -module msp ../idl/rpc/oss.thrift
mkdir -p oss/internal/infra/mysql/model
cd oss
kitex -service oss -module msp -use "msp/biz_server/kitex_gen" ../../idl/rpc/oss.thrift

cd internal/infra/mysql/model
go run entgo.io/ent/cmd/ent new DomainMapping
go run entgo.io/ent/cmd/ent new Tenant
# 需创建数据库模型，修改schema

go mod tidy
go-mod-upgrade
```

## 整体架构

### gateway

- 通过 hertz 作为纯网关，泛化调用下游业务服务
- consul 作为注册中心
- hertz 遍历 idl 获取泛化调用服务，暂时无热加载功能，后续尝试迭代

### biz_server

- kitex 分两种业务服务
- 服务编排，调用各基础服务数据组装
- 基础服务，仅提供基础功能的基建服务

## 网关配置

```json
{
  "name": "gateway",
  "port": 8080,
  "gateway_resource": [
    {
      "route": "mtop/qiniu/upload",
      "svr_name": "oss",
      "finger_print": "/cur/upload",
      "parent_path": "your idl parent path，idl_path所在文件夹的绝对路径",
      "idl_path": "oss.thrift",
      "include_path": [
        "../base/common.thrift"
      ]
    }
  ]
}
```

## 实现方式

- 通过 consul 中维护的 gateway_resource 定位到具体的接口
- route 为网关暴露的 http 路由
- 通过路由，匹配到对应的 finger_print 和 provider
- 调用 provider 下的 finger_print 方法
- finger_print 方法为 idl 中 service 下的 api.post
- 引入 finger_print 是为了后续灵活处理
    - 多个 route 映射到相同泛化调用方法
    - 同时后续拆分、升级外部可以无感知（调用 http 方法不变，底层微服务变更）
- 热加载本身不复杂，核心是 svrRouteMap 修改（未做）
    - 配置中心触发gateway中的配置更新，删除无效数据，加载新数据到 svrRouteMap
    - idl替换需要有配置管理中心或者将文件写入数据库中

## 后续计划

- 修改为 Multiple Services 模式。早期kitex仅支持单 server 单 service，现在支持多service了，升级
- 支持 Reflection 后，跟进升级为不依赖idl的泛化调用
