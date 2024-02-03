package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"msp/biz_server/kitex_gen/oss"
	"msp/biz_server/oss/internal/infra/qiniu"
	"msp/biz_server/oss/internal/usecase"
	"msp/common/model/errors"
	"os"
	"path"
)

type OssUploadService struct {
	ctx           context.Context
	ossRepository usecase.OssRepository
	repository    usecase.Repository
}

func NewOssUploadService(ctx context.Context, repository usecase.Repository) *OssUploadService {
	uploader := qiniu.NewOssUploader()
	uploader.WithContext(ctx)
	return &OssUploadService{
		ctx:           ctx,
		ossRepository: uploader,
		repository:    repository,
	}
}

func (s *OssUploadService) Run(req *oss.UploadRequest) (resp *oss.OssUploadResp, err error) {
	klog.CtxInfof(s.ctx, "OssUploadService Run req: %+v", req)
	err = req.IsValid()
	if err != nil {
		klog.CtxWarnf(s.ctx, "OssUploadService Run req valid err %+v", err)
		err = errors.NewErrNo(errors.ParamErrCode, err.Error())
		return
	}
	_, err = os.Stat(req.FileName)
	if err != nil {
		klog.CtxErrorf(s.ctx, "OssUploadService Run file stat err: %+v", err)
		err = errors.NewErrNo(errors.ServiceErrCode, err.Error())
		return
	}
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

	remote := path.Join(req.GetRemoteDir(), req.GetRemoteName())
	if !req.GetForceUpload() {
		if s.ossRepository.CheckStat(tenant.AccessKey, tenant.SecretKey, mapping.RegionID, mapping.BucketName, remote) {
			err = errors.NewErrNo(errors.ServiceErrCode, "file already exists")
			klog.CtxWarnf(s.ctx, "OssUploadService Run tenant err: %+v", err)
			return
		}
	}

	err = s.ossRepository.Upload(tenant.AccessKey, tenant.SecretKey, mapping.RegionID, mapping.BucketName, req.GetFileName(), remote)
	if err != nil {
		klog.CtxErrorf(s.ctx, "OssUploadService Run ossRepository Upload err: %+v", err)
		err = errors.NewErrNo(errors.ServiceErrCode, err.Error())
		return
	}
	resp = new(oss.OssUploadResp)
	resp.RemoteUrl = path.Join(mapping.Domain, remote)
	return
}
