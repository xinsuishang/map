package usecase

import (
	"context"
	"msp/biz_server/gpt/internal/entity"
)

// Repository 资源操作接口定义
type Repository interface {
	GetTenantsById(ctx context.Context, id int32) (*entity.TenantEntity, error)
	GetCacheTenantsAndParentById(ctx context.Context, id int32) (*entity.TenantEntity, *entity.TenantEntity, error)
}
