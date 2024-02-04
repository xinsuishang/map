package model

import (
	"context"
	"errors"
	"github.com/bytedance/gopkg/cache/asynccache"
	"msp/biz_server/gpt/internal/entity"
	"msp/biz_server/gpt/internal/infra/mysql/model/ent"
	"msp/biz_server/gpt/internal/infra/mysql/model/ent/chat"
	"time"
)

var chatEntityCache = asynccache.NewAsyncCache(asynccache.Options{
	// 1分钟过期，保证不自动刷新
	RefreshDuration: time.Minute * 1 * 2,
	EnableExpire:    true,
	ExpireDuration:  time.Minute * 1,
	Fetcher: func(sessionId string) (interface{}, error) {
		getChat, err := DB.GetChat(context.Background(), sessionId)
		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		}
		return getChat, nil
	},
})

func getFromChatCache(sessionId string) (*entity.ChatEntity, error) {
	val, err := chatEntityCache.Get(sessionId)
	if val == nil || err != nil {
		return nil, err
	}

	chatEntity, ok := val.(*entity.ChatEntity)
	if ok {
		return chatEntity, err
	}

	return nil, errors.New("error assert")
}

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
