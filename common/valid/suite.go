package valid

import (
	"context"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/pkg/utils/kitexutil"
	"github.com/cloudwego/kitex/server"
)

type validSuite struct {
	sOpts []server.Option
}

func NewServerSuite(opt ...server.Option) *validSuite {
	suite := &validSuite{}
	suite.Options()
	return suite
}
func (s *validSuite) Options() []server.Option {
	return []server.Option{
		server.WithMiddleware(middleware),
	}
}

func middleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
		if arg, ok := request.(utils.KitexArgs); ok {
			if req, ok := arg.GetFirstArgument().(ValidIface); ok && req != nil {
				err := Validate(req)
				if err != nil {
					rpcInfo, ok := kitexutil.GetRPCInfo(ctx)
					var serviceName, methodName string
					if ok {
						serviceName = rpcInfo.Invocation().ServiceName()
						methodName = rpcInfo.Invocation().MethodName()
					}
					klog.CtxErrorf(ctx, "service: %s method: %s Validate failed: %v", serviceName, methodName, err)
					return err
				}
			}
		}
		err := next(ctx, request, response)
		//if result, ok := response.(utils.KitexResult); ok {
		//	if resp, ok := result.GetResult().(*echo.Response); ok {
		//		klog.Debugf("Response Message: %v", resp.Message)
		//		// resp.SetSuccess(...) 可以用于替换自定义的响应结果
		//		// 但要注意：类型应与该 method 的结果类型相同
		//	}
		//}
		return err
	}
}
