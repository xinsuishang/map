package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"msp/biz_server/gpt/internal/entity"
	"msp/biz_server/gpt/internal/usecase"
	"msp/biz_server/kitex_gen/gpt"
	"msp/common/model/errors"
)

type ChatApplicationService struct {
	ctx        context.Context
	repository usecase.Repository
}

func NewChatApplicationService(ctx context.Context, repository usecase.Repository) *ChatApplicationService {
	return &ChatApplicationService{
		ctx:        ctx,
		repository: repository,
	}
}

func (s *ChatApplicationService) Run(req *gpt.ApplicationListReq) (resp *gpt.ApplicationListResp, err error) {
	klog.CtxInfof(s.ctx, "ChatApplicationService Run req: %+v", req)
	resp = gpt.NewApplicationListResp()
	resp.SetPage(req.GetPage())
	list, total, err := s.repository.GetTenantList(s.ctx, req.GetModelId(), req.GetParentId(), req.GetPage().GetPageNo(), req.GetPage().GetPageSize())
	if err != nil {
		klog.CtxErrorf(s.ctx, "ChatApplicationService Run GetList err: %+v", err)
		err = errors.NewErrNo(errors.ServiceErrCode, err.Error())
		return
	}
	resp.GetPage().SetTotal(total)
	resp.SetList(convertList(list))
	return
}

func convert(source *entity.TenantEntity) *gpt.ApplicationInfo {
	return &gpt.ApplicationInfo{
		ModelId:       source.ID,
		ParentId:      source.ParentId,
		Name:          source.Name,
		Model:         source.Model,
		IsApplication: source.IsApplication,
		Desc:          source.Desc,
		Dashboard:     source.Dashboard,
	}
}

func convertList(source []*entity.TenantEntity) []*gpt.ApplicationInfo {
	if source == nil {
		return make([]*gpt.ApplicationInfo, 0)
	}
	ret := make([]*gpt.ApplicationInfo, 0, len(source))
	for i := range source {
		ret = append(ret, convert(source[i]))
	}
	return ret
}
