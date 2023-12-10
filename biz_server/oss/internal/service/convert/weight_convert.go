package convert

import (
	"msp/biz_server/kitex_gen/oss"
)

func WeightToOssReq(source *oss.WeightUploadRequest, remoteDir string, domainId int32) (target *oss.UploadRequest) {
	target = new(oss.UploadRequest)
	target.SetFileName(source.GetFileName())
	target.SetForceUpload(source.GetForceUpload())
	target.SetRemoteName(source.GetRemoteName())
	target.SetRemoteDir(remoteDir)
	target.SetDomainId(domainId)
	return
}

func WeightToNotionReq(source *oss.WeightUploadRequest, remoteUrl string, domainId int32) (target *oss.NotionUploadRequest) {
	target = new(oss.NotionUploadRequest)
	target.SetWeight(source.GetWeight())
	target.SetDomainId(4)
	target.SetDateTime(source.GetDateTime())
	target.SetFileUrl(remoteUrl)
	target.SetDomainId(domainId)
	return
}
