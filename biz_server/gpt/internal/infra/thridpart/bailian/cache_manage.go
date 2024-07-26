package bailian

import (
	"context"
	"errors"
	client "github.com/aliyun/alibabacloud-bailian-go-sdk/client"
	"github.com/bytedance/gopkg/cache/asynccache"
	"github.com/cloudwego/kitex/pkg/klog"
	"msp/biz_server/gpt/internal/infra/mysql/model"
	"strconv"
	"time"
)

type accessTokenClient struct {
	client.AccessTokenClient
	AppId     string
	IsDeleted bool
}

func deleteCacheIf() {
	for {
		time.Sleep(time.Second * 5)
		tokenClientCache.DeleteIf(func(modelId string) bool {
			primary, err := strconv.Atoi(modelId)

			if err != nil {
				return true
			}
			token, err := getToken(int32(primary))
			if err != nil {
				return true
			}
			return token.IsDeleted
		})
	}
}

func convertClient(source interface{}) (*accessTokenClient, error) {
	// 根据client获取token
	tokenClient, ok := source.(*accessTokenClient)
	if ok {
		return tokenClient, nil
	}
	// token 本身 sdk 提供了内存缓存逻辑
	return nil, errors.New("error assert")
}

var tokenClientCache = asynccache.NewAsyncCache(asynccache.Options{
	RefreshDuration: time.Hour * 6,
	IsSame: func(key string, oldData, newData interface{}) bool {
		klog.Info("tokenClientCache refresh")
		return oldData == newData
	},
	Fetcher: func(modelId string) (interface{}, error) {
		primary, err := strconv.Atoi(modelId)
		if err != nil {
			return nil, err
		}
		entityData, parentEntity, err := model.DB.GetCacheTenantAndParentById(context.Background(), int32(primary))

		// 根据 app 参数获取 tokenClient
		cacheData := &accessTokenClient{
			AppId: entityData.AccessKey,
		}
		cacheData.AccessKeyId = parentEntity.AccessKey
		cacheData.AccessKeySecret = parentEntity.SecretKey
		cacheData.AgentKey = entityData.SecretKey
		cacheData.IsDeleted = entityData.IsDeleted || parentEntity.IsDeleted
		return cacheData, nil
	},
})

func getToken(modelId int32) (*accessTokenClient, error) {
	// 根据 key 获取clientCache
	val, err := tokenClientCache.Get(strconv.Itoa(int(modelId)))

	if err != nil {
		return nil, err
	}
	return convertClient(val)
}
