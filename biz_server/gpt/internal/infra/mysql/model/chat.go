package model

import (
	"context"
	"errors"
	"msp/biz_server/gpt/internal/entity"
	"msp/biz_server/gpt/internal/infra/mysql/model/ent"
	"msp/biz_server/gpt/internal/infra/mysql/model/ent/chat"
)

func (p *repository) GetChat(ctx context.Context, sessionId string) (*entity.ChatEntity, error) {
	only, err := p.db.Chat.Query().Where(chat.SessionID(sessionId)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return convertChat(only), nil
}

func (p *repository) GetChatCache(ctx context.Context, sessionId string) (*entity.ChatEntity, error) {
	return getFromChatCache(sessionId)
}

func (p *repository) SaveChat(ctx context.Context, modelId int32, sessionId, name string) (*entity.ChatEntity, error) {
	cache, err := p.GetChatCache(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	if cache == nil {
		create := p.db.Chat.Create().
			SetModelID(modelId).
			SetSessionID(sessionId).
			SetName(name)
		save, err := create.Save(ctx)
		if err != nil {
			return nil, err
		}
		return convertChat(save), err
	}

	if cache.ModelID != modelId {
		return nil, errors.New("modelId sessionId not match")
	}

	return cache, err
}

func convertChat(source *ent.Chat) *entity.ChatEntity {
	return &entity.ChatEntity{
		ID:        source.ID,
		ModelID:   source.ModelID,
		SessionID: source.SessionID,
		Name:      source.Name,
	}
}
