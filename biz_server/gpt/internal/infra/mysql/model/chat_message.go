package model

import (
	"context"
	"msp/biz_server/gpt/internal/entity"
	"msp/biz_server/gpt/internal/infra/mysql/model/ent"
)

func (p *repository) SaveMessage(ctx context.Context, chatId int32, version int8, requestId, text string) (*entity.ChatMessageEntity, error) {
	create := p.db.ChatMessage.Create().
		SetChatID(chatId).
		SetRequestID(requestId).
		SetText(text).
		SetVersion(version)
	save, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertMessage(save), err
}

func convertMessage(source *ent.ChatMessage) *entity.ChatMessageEntity {
	return &entity.ChatMessageEntity{
		ID:        source.ID,
		ChatID:    source.ChatID,
		RequestID: source.RequestID,
		Text:      source.Text,
		CreatedAt: source.CreatedAt,
		UpdatedAt: source.UpdatedAt}
}
