package thridpart

import (
	"context"
	"errors"
	"msp/biz_server/gpt/internal/infra/mysql/model"
	"msp/biz_server/gpt/internal/infra/thridpart/bailian"
)

const (
	BaiLian = "bailian" //百炼
)

func Dispatch(ctx context.Context, modelId int32, requestId, sessionId, prompt string) (string, error) {
	tenantEntity, err := model.DB.GetCacheTenantById(ctx, modelId)
	if err != nil {
		return "", err
	}

	if tenantEntity == nil || tenantEntity.IsDeleted {
		return "", errors.New("tenant not found or deleted")
	}

	curModel := tenantEntity.Model
	switch curModel {
	case BaiLian:
		return bailian.CreateCompletion(ctx, modelId, requestId, sessionId, prompt)
	default:
		return "", errors.New("model unsupported")
	}
}
