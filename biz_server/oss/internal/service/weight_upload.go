package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"msp/biz_server/kitex_gen/common"
	"msp/biz_server/kitex_gen/oss"
	"msp/biz_server/oss/internal/service/convert"
	"msp/biz_server/oss/internal/usecase"
	"msp/common/model/errors"
)

type WeightUploadService struct {
	ctx        context.Context
	repository usecase.Repository
}

func NewWeightUploadService(ctx context.Context, repository usecase.Repository) *WeightUploadService {
	return &WeightUploadService{ctx: ctx,
		repository: repository,
	}
}

func (s *WeightUploadService) Run(req *oss.WeightUploadRequest) (resp *common.EmptyResponse, err error) {
	klog.CtxInfof(s.ctx, "WeightUploadService Run req: %+v", req)
	uploadService := NewOssUploadService(s.ctx, s.repository)
	uploadResp, err := uploadService.Run(convert.WeightToOssReq(req, 1))
	if err != nil {
		err = errors.NewErrNo(errors.ServiceErrCode, err.Error())
		return
	}

	remoteUrl := uploadResp.GetRemoteUrl()

	notionService := NewNotionUploadService(s.ctx, s.repository)
	return notionService.Run(convert.WeightToNotionReq(req, remoteUrl, 4))
}
