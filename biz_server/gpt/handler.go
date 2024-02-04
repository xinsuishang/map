package main

import (
	"context"
	"msp/biz_server/gpt/internal/service"
	"msp/biz_server/gpt/internal/usecase"
	gpt "msp/biz_server/kitex_gen/gpt"
)

// ChatServiceImpl implements the last service interface defined in the IDL.
type ChatServiceImpl struct {
	repository usecase.Repository
}

// Chat implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) Chat(ctx context.Context, req *gpt.ApplicationReq) (resp *gpt.ApplicationResp, err error) {
	return service.NewChatService(ctx, s.repository).Run(req)
}

// ChatApplicationList implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) ChatApplicationList(ctx context.Context, req *gpt.ApplicationListReq) (resp *gpt.ApplicationListResp, err error) {
	return service.NewChatApplicationService(ctx, s.repository).Run(req)
}
