package qiniu

import (
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/auth/qbox"

	"github.com/qiniu/go-sdk/v7/storage"
)

func authOss(accessKey, secretKey string) *auth.Credentials {
	return qbox.NewMac(accessKey, secretKey)
}

func config(regionId string) *storage.Config {
	id, b := storage.GetRegionByID(storage.RegionID(regionId))

	if !b {
		panic(nil)
	}
	return &storage.Config{
		UseHTTPS: false, // 是否使用 HTTPS
		Region:   &id,
	}
}
