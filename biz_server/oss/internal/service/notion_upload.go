package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"msp/biz_server/kitex_gen/common"
	"msp/biz_server/kitex_gen/oss"
	"msp/biz_server/oss/internal/infra/notion"
	"msp/biz_server/oss/internal/infra/notion/model"
	"msp/biz_server/oss/internal/usecase"
	"msp/common/model/errors"
)

type NotionUploadService struct {
	ctx              context.Context
	notionRepository usecase.NotionRepository
	repository       usecase.Repository
}

func NewNotionUploadService(ctx context.Context, repository usecase.Repository) *NotionUploadService {
	database := notion.NewNotionDatabase()
	database.WithContext(ctx)
	return &NotionUploadService{
		ctx:              ctx,
		notionRepository: database,
		repository:       repository,
	}
}

func (s *NotionUploadService) Run(req *oss.NotionUploadRequest) (resp *common.EmptyResponse, err error) {
	klog.CtxInfof(s.ctx, "NotionUploadService Run req: %+v", req)
	resp = common.NewEmptyResponse()
	mapping, err := s.repository.GetDomainMapping(s.ctx, req.GetDomainId())
	if err != nil {
		klog.CtxErrorf(s.ctx, "OssUploadService Run mapping err: %+v", err)
		err = errors.NewErrNo(errors.ServiceErrCode, err.Error())
		return
	}
	tenant, err := s.repository.GetTenantsById(s.ctx, mapping.TenantID)
	if err != nil {
		klog.CtxErrorf(s.ctx, "OssUploadService Run tenant err: %+v", err)
		err = errors.NewErrNo(errors.ServiceErrCode, err.Error())
		return
	}
	dataMap := make(map[string]model.PropertyData)
	dataMap["date"] = *model.NewDate(req.GetDateTime())
	dataMap["证明"] = *model.NewUrl(req.GetFileUrl())
	dataMap["当日体重"] = *model.NewNumber(req.GetWeight())
	dataMap["Files & media"] = *model.NewFiles(req.GetFileUrl())
	dataMap["批次"] = *model.NewTitles("2023年2月减肥")

	err = s.notionRepository.AddPageToDatabase(tenant.SecretKey, mapping.BucketName, dataMap)
	if err != nil {
		klog.CtxErrorf(s.ctx, "OssUploadService Run AddPageToDatabase err: %+v", err)
		err = errors.NewErrNo(errors.ServiceErrCode, err.Error())
	}
	return
}
