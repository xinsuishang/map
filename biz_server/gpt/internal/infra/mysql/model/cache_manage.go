package model

import (
	"context"
	"errors"
	"github.com/bytedance/gopkg/cache/asynccache"
	"msp/biz_server/gpt/internal/entity"
	"msp/biz_server/gpt/internal/infra/mysql/model/ent"
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
