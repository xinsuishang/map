package model

import (
	"context"
	"msp/biz_server/oss/internal/entity"
	"msp/biz_server/oss/internal/infra/mysql/model/ent"
	"msp/biz_server/oss/internal/usecase"
)

type repository struct {
	db *ent.Client
}

func NewEntClient(db *ent.Client) usecase.Repository {
	return &repository{db: db}
}

func (p *repository) GetTenantsById(ctx context.Context, id int32) (*entity.TenantEntity, error) {
	get, err := p.db.Tenant.Get(ctx, int(id))
	if err != nil {
		return nil, err
	}
	return &entity.TenantEntity{
		ID:        int32(get.ID),
		Name:      get.Name,
		Region:    get.Region,
		Type:      get.Type,
		AccessKey: get.AccessKey,
		SecretKey: get.SecretKey,
		Desc:      get.Desc,
	}, nil
}
func (p *repository) GetDomainMapping(ctx context.Context, id int32) (*entity.DomainMappingEntity, error) {
	get, err := p.db.DomainMapping.Get(ctx, int(id))
	if err != nil {
		return nil, err
	}
	return &entity.DomainMappingEntity{
		ID:         int32(get.ID),
		TenantID:   int32(get.TenantID),
		RegionID:   get.RegionID,
		Domain:     get.Domain,
		BucketName: get.BucketName,
		Desc:       get.Desc,
	}, nil
}
