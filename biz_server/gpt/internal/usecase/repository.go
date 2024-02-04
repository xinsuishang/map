package usecase

import (
	"context"
	"msp/biz_server/gpt/internal/entity"
)

// Repository 资源操作接口定义
type Repository interface {
	GetTenantsById(ctx context.Context, id int32) (*entity.TenantEntity, error)
	GetCacheTenantsAndParentById(ctx context.Context, id int32) (*entity.TenantEntity, *entity.TenantEntity, error)
	GetTenantList(ctx context.Context, modelId, parentId, pageNo, pageSize int32) ([]*entity.TenantEntity, int32, error)

	GetChatCache(ctx context.Context, sessionId string) (*entity.ChatEntity, error)
	SaveChat(ctx context.Context, modelId int32, sessionId, name string) (*entity.ChatEntity, error)

	SaveMessage(ctx context.Context, chatId int32, version int8, requestId, text string) (*entity.ChatMessageEntity, error)
}
