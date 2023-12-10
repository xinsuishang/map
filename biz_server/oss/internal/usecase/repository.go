package usecase

import (
	"context"
	"msp/biz_server/oss/internal/entity"
)

// Repository 资源操作接口定义
type Repository interface {
	GetTenantsById(ctx context.Context, id int32) (*entity.TenantEntity, error)
	GetDomainMapping(ctx context.Context, id int32) (*entity.DomainMappingEntity, error)
}
