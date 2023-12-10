package usecase

import "context"

type OssRepository interface {
	WithContext(ctx context.Context)
	Upload(accessKey, secretKey, regionId, bucketName, localFile, remoteFile string) (err error)
	CheckStat(accessKey, secretKey, regionId, bucketName, remoteFile string) (res bool)
}
