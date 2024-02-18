package convert

import (
	"msp/biz_server/kitex_gen/oss"
	"path"
)

func WeightToOssReq(source *oss.WeightUploadRequest, domainId int32) (target *oss.UploadRequest) {
	target = new(oss.UploadRequest)
	target.SetFileName(source.GetFileName())
	target.SetForceUpload(source.GetForceUpload())
	dir, file := path.Split(source.GetRemoteName())
	target.SetRemoteDir(dir)
	target.SetRemoteName(file)
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
	target.SetTitle(source.GetTitle())
	return
}
