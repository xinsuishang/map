package service

import (
	"context"
	"msp/biz_server/gpt/internal/usecase"
)

type ChatService struct {
	ctx        context.Context
	repository usecase.Repository
}

func NewChatService(ctx context.Context, repository usecase.Repository) *ChatService {
	return &ChatService{
		ctx:        ctx,
		repository: repository,
	}
}
