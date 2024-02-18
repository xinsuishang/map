package bailian

import (
	"context"
	"errors"
	client "github.com/aliyun/alibabacloud-bailian-go-sdk/client"
	"github.com/cloudwego/kitex/pkg/klog"
)

func CreateCompletion(ctx context.Context, modelId int32, requestId, sessionId, prompt string) (string, error) {
	tokenClient, err := getToken(modelId)
	if err != nil {
		return "", err
	}
	if tokenClient.IsDeleted {
		return "", errors.New("application is deleted")
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
