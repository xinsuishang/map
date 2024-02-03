package bailian

import (
	"errors"
	client "github.com/aliyun/alibabacloud-bailian-go-sdk/client"
	"github.com/bytedance/gopkg/cache/asynccache"
	"github.com/cloudwego/kitex/pkg/klog"
	"msp/biz_server/gpt/internal/entity"
	"time"
)

type accessTokenClient struct {
	client.AccessTokenClient
	AppId string
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
	RefreshDuration: time.Minute * 1,
	EnableExpire:    true,
	ExpireDuration:  time.Minute * 1 * 2,
	IsSame: func(key string, oldData, newData interface{}) bool {
		cacheData, _ := convertClient(oldData)
		generateData, _ := convertClient(newData)

		if cacheData.AccessKeyId != generateData.AccessKeyId ||
			cacheData.AccessKeySecret != generateData.AccessKeySecret ||
			cacheData.AgentKey != generateData.AgentKey {
			// 新client，则先初始化token
			_, _ = generateData.GetToken()
			return false
		}
		// 缓存的结果是个client，内部会携带token缓存，生成成本较高
		// 强行替换缓存结果，如果token即将过期，则主动替换
		data := cacheData.TokenData
		unix := time.Now().Unix()
		if (*data.ExpiredTime - unix) < (600 + 90) {
			go func() {
				result, err := cacheData.CreateToken()
				if err != nil {
					return
				}
				data.SetToken(*result.Token)
				data.SetExpiredTime(*result.ExpiredTime)
			}()
		}
		newData = oldData
		return true
	},
	Fetcher: func(modelId string) (interface{}, error) {

		// todo 根据 modelId 获取 app 参数
		entityData := &entity.TenantEntity{}
		parentEntity := &entity.TenantEntity{}

		// 根据 app 参数获取 tokenClient
		cacheData := &accessTokenClient{
			AppId: entityData.AccessKey,
		}
		cacheData.AccessKeyId = parentEntity.AccessKey
		cacheData.AccessKeySecret = parentEntity.SecretKey
		cacheData.AgentKey = entityData.SecretKey
		return cacheData, nil
	},
})

func getToken(modelId int32) (*accessTokenClient, error) {
	// 根据 key 获取clientCache
	val, err := tokenClientCache.Get(string(modelId))

	if err != nil {
		return nil, err
	}
	return convertClient(val)
}

func CreateCompletion(modelId int32, requestId, sessionId, prompt string) (string, error) {
	tokenClient, err := getToken(modelId)
	if err != nil {
		return "", err
	}
	token, err := tokenClient.GetToken()
	klog.Infof("expiretTime %d", tokenClient.TokenData.ExpiredTime)
	if err != nil {
		return "", err
	}
	cc := client.CompletionClient{Token: token}
	request := &client.CompletionRequest{}
	request.AppId = tokenClient.AppId
	request.Prompt = prompt
	request.SessionId = sessionId
	request.RequestId = requestId
	response, err := cc.CreateCompletion(request)
	if err != nil {
		klog.Error("%v\n", err.Error())
		return "", err
	}

	if !response.Success {
		klog.Error("failed to create completion, requestId: %s, code: %s, message: %s\n",
			response.GoString())
		return "", errors.New(response.Message)
	}

	if response.Data.Usage != nil && len(response.Data.Usage) > 0 {
		usage := response.Data.Usage[0]
		klog.Infof(", inputTokens: %d, outputTokens: %d\n", usage.InputTokens, usage.OutputTokens)
	}
	klog.Infof("requestId: %s, text : %s", response.RequestId, response.Data.Text)
	return response.Data.Text, nil
}
