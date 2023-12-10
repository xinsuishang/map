package notion

import (
	"context"
	"msp/biz_server/oss/internal/infra/notion/model"
	"msp/biz_server/oss/internal/usecase"
)

type notionDatabase struct {
	ctx context.Context
}

func NewNotionDatabase() usecase.NotionRepository {
	return &notionDatabase{}
}
func (p *notionDatabase) WithContext(ctx context.Context) {
	p.ctx = ctx
}

func (p *notionDatabase) AddPageToDatabase(secret, databaseId string, dataMap map[string]model.PropertyData) error {
	holder := NewNotionHolder(secret)

	data := &model.CommonData{
		Properties: dataMap,
	}

	if len(databaseId) > 0 {
		parentData := &model.PropertyData{
			DatabaseId: databaseId,
		}
		parentData.Type = "database_id"
		data.Parent = parentData
	}
	return holder.Do(data)
}
