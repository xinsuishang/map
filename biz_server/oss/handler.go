package main

import (
	"context"
	"msp/biz_server/kitex_gen/common"
	"msp/biz_server/kitex_gen/oss"
	"msp/biz_server/oss/internal/service"
	"msp/biz_server/oss/internal/usecase"
)

// UploadServiceImpl implements the last service interface defined in the IDL.
type UploadServiceImpl struct {
	repository usecase.Repository
}

// OssUpload implements the UploadServiceImpl interface.
func (s *UploadServiceImpl) OssUpload(ctx context.Context, req *oss.UploadRequest) (resp *oss.OssUploadResp, err error) {
	return service.NewOssUploadService(ctx, s.repository).Run(req)
}

// NotionUpload implements the UploadServiceImpl interface.
func (s *UploadServiceImpl) NotionUpload(ctx context.Context, req *oss.NotionUploadRequest) (resp *common.EmptyResponse, err error) {
	return service.NewNotionUploadService(ctx, s.repository).Run(req)
}

// WeightUpload implements the UploadServiceImpl interface.
func (s *UploadServiceImpl) WeightUpload(ctx context.Context, req *oss.WeightUploadRequest) (resp *common.EmptyResponse, err error) {
	return service.NewWeightUploadService(ctx, s.repository).Run(req)
}
