package usecase

import (
	"context"
	"msp/biz_server/oss/internal/infra/notion/model"
)

type NotionRepository interface {
	WithContext(ctx context.Context)
	AddPageToDatabase(secret, databaseId string, dataMap map[string]model.PropertyData) error
}
