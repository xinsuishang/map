package handler

import (
	"bytes"
	"context"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"msp/common/model/errors"
	"msp/gateway/biz/model"
	"net/http"
)

var SvcRouteMap = make(map[string]genericclient.Client)
var SvcFingerPrintMap = make(map[string]string)

// Gateway handler
func Gateway(ctx context.Context, c *app.RequestContext) {
	resp := new(model.Response)
	defer c.JSON(http.StatusOK, resp)

	routeName := c.Param("path")
	cli, ok := SvcRouteMap[routeName]
	if !ok {
		resp.Err(errors.ServiceNotFoundErr.WithMessage("404"))
		return
	}
	var params model.RequiredParams
	if err := c.BindAndValidate(&params); err != nil {
		hlog.CtxErrorf(ctx, "%v", err)
		resp.Err(errors.ParamErr.WithError(err))
		return
	}
	body, err := sonic.Marshal(params.Body)
	if err != nil {
		hlog.CtxErrorf(ctx, "%v", err)
		resp.Err(errors.ParamErr.WithError(err))
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "", bytes.NewBuffer(body))
	if err != nil {
		hlog.CtxWarnf(ctx, "new http request failed: %v", err)
		resp.Err(errors.ServerHandleFail)
		return
	}
	req.URL.Path = SvcFingerPrintMap[routeName]

	customReq, err := generic.FromHTTPRequest(req)
	if err != nil {
		hlog.CtxErrorf(ctx, "convert request failed: %v", err)
		resp.Err(errors.ServerHandleFail)
		return
	}
	remoteResp, err := cli.GenericCall(ctx, "", customReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "GenericCall err:%v", err)
		bizErr, ok := kerrors.FromBizStatusError(err)
		if !ok {
			resp.BizErrMsg(err.Error())
			return
		}
		resp.Code = bizErr.BizStatusCode()
		resp.Message = bizErr.BizMessage()
		return
	}
	realResp, ok := remoteResp.(*generic.HTTPResponse)
	if !ok {
		hlog.CtxErrorf(ctx, "remoteResp err:%v", err)
		resp.BizErrMsg(err.Error())
		return
	}

	resp.Success(realResp.Body)
}

func GetawayList(ctx context.Context, c *app.RequestContext) {
	resp := new(model.Response)
	defer c.JSON(http.StatusOK, resp)

	resp.Success(SvcFingerPrintMap)
}
