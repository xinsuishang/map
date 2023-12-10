package qiniu

import (
	"context"
	"github.com/qiniu/go-sdk/v7/storage"
	"msp/biz_server/oss/internal/usecase"
)

type ossUploader struct {
	ctx context.Context
}

func NewOssUploader() usecase.OssRepository {
	return &ossUploader{}
}

func (p *ossUploader) WithContext(ctx context.Context) {
	p.ctx = ctx
}
func (p *ossUploader) Upload(accessKey, secretKey, regionId, bucketName, localFile, remoteFile string) (err error) {
	mac := authOss(accessKey, secretKey)
	cfg := config(regionId)

	// 创建七牛云上传管理对象
	formUploader := storage.NewFormUploader(cfg)
	putPolicy := storage.PutPolicy{
		Scope: bucketName,
	}
	upToken := putPolicy.UploadToken(mac)

	// 创建上传的文件
	ret := storage.PutRet{}

	// 上传文件
	err = formUploader.PutFile(context.Background(), &ret, upToken, remoteFile, localFile, nil)
	return
}
func (p *ossUploader) CheckStat(accessKey, secretKey, regionId, bucketName, remoteFile string) (res bool) {
	mac := authOss(accessKey, secretKey)
	cfg := config(regionId)
	bucketManager := storage.NewBucketManager(mac, cfg)
	_, err := bucketManager.Stat(bucketName, remoteFile)
	return err == nil
}
