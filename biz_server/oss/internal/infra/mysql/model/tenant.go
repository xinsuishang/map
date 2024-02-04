package model

import (
	"context"
	"msp/biz_server/oss/internal/entity"
)

func (p *repository) GetTenantsById(ctx context.Context, id int32) (*entity.TenantEntity, error) {
	get, err := p.db.Tenant.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &entity.TenantEntity{
		ID:        get.ID,
		Name:      get.Name,
		Region:    get.Region,
		Type:      get.Type,
		AccessKey: get.AccessKey,
		SecretKey: get.SecretKey,
		Desc:      get.Desc,
	}, nil
}
func (p *repository) GetDomainMapping(ctx context.Context, id int32) (*entity.DomainMappingEntity, error) {
	get, err := p.db.DomainMapping.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &entity.DomainMappingEntity{
		ID:         get.ID,
		TenantID:   get.TenantID,
		RegionID:   get.RegionID,
		Domain:     get.Domain,
		BucketName: get.BucketName,
		Desc:       get.Desc,
	}, nil
}
