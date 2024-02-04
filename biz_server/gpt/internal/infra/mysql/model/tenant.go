package model

import (
	"context"
	"errors"
	"github.com/bytedance/gopkg/cache/asynccache"
	"msp/biz_server/gpt/internal/entity"
	"msp/biz_server/gpt/internal/infra/mysql/model/ent"
	"msp/biz_server/gpt/internal/infra/mysql/model/ent/tenant"
	"msp/common/simpleutils"
	"strconv"
	"time"
)

var tenantEntityCache = asynccache.NewAsyncCache(asynccache.Options{
	// 1分钟过期，保证不自动刷新
	RefreshDuration: time.Minute * 1 * 2,
	EnableExpire:    true,
	ExpireDuration:  time.Minute * 1,
	Fetcher: func(id string) (interface{}, error) {
		if id == "-1" {
			return nil, nil
		}
		primary, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		return DB.GetTenantsById(context.Background(), int32(primary))
	},
})

func getFromTenantCache(id int32) (*entity.TenantEntity, error) {
	val, err := tenantEntityCache.Get(strconv.Itoa(int(id)))
	if val == nil || err != nil {
		return nil, err
	}

	tenantEntity, ok := val.(*entity.TenantEntity)
	if ok {
		return tenantEntity, err
	}

	return nil, errors.New("error assert")
}

func (p *repository) GetTenantsById(ctx context.Context, id int32) (*entity.TenantEntity, error) {
	get, err := p.db.Tenant.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return convertTenant(get), nil
}
func (p *repository) GetCacheTenantsAndParentById(ctx context.Context, id int32) (*entity.TenantEntity, *entity.TenantEntity, error) {

	cacheEntity, err := getFromTenantCache(id)
	if cacheEntity == nil || err != nil {
		return cacheEntity, nil, err
	}
	parentId := cacheEntity.ParentId
	parentEntity, err := getFromTenantCache(parentId)
	return cacheEntity, parentEntity, err
}

func (p *repository) GetTenantList(ctx context.Context, modelId, parentId, pageNo, pageSize int32) ([]*entity.TenantEntity, int32, error) {
	query := p.db.Tenant.Query()
	if modelId != 0 {
		query.Where(tenant.ID(modelId))
	}
	if parentId != 0 {
		query.Where(tenant.ParentID(parentId))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	total := int32(count)
	page := simpleutils.CalcTotalPage(pageSize, total)
	if page < pageNo {
		return make([]*entity.TenantEntity, 0), 0, err
	}
	all, err := query.
		Limit(int(pageSize)).
		Offset(int(pageSize * (pageNo - 1))).
		All(ctx)
	return convertTenantList(all), total, err
}

func convertTenant(source *ent.Tenant) *entity.TenantEntity {
	return &entity.TenantEntity{
		ID:            source.ID,
		ParentId:      source.ParentID,
		Name:          source.Name,
		Model:         source.Model,
		IsApplication: source.IsApplication,
		AccessKey:     source.AccessKey,
		SecretKey:     source.SecretKey,
		Desc:          source.Desc,
		Dashboard:     source.Dashboard,
	}
}

func convertTenantList(source []*ent.Tenant) []*entity.TenantEntity {
	if source == nil {
		return make([]*entity.TenantEntity, 0)
	}
	target := make([]*entity.TenantEntity, 0, len(source))
	for i := range source {
		target = append(target, convertTenant(source[i]))
	}
	return target
}
