package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"msp/biz_server/gpt/internal/infra/thridpart/bailian"
	"msp/biz_server/gpt/internal/usecase"
	"msp/biz_server/kitex_gen/gpt"
	"msp/common/model/errors"
	"strings"
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

func (s *ChatService) Run(req *gpt.ApplicationReq) (resp *gpt.ApplicationResp, err error) {
	klog.CtxInfof(s.ctx, "ChatService Run req: %+v", req)
	resp = gpt.NewApplicationResp()

	if req.GetSessionId() == "" {
		req.SetSessionId(strings.ReplaceAll(uuid.New().String(), "-", ""))
	}
	resp.SetModelId(req.GetModelId())
	resp.SetSessionId(req.GetSessionId())

	chat, err := s.repository.SaveChat(s.ctx, req.GetModelId(), req.GetSessionId(), req.GetPrompt())

	if err != nil {
		klog.CtxErrorf(s.ctx, "ChatService Run saveChat err: %+v", err)
		err = errors.NewErrNo(errors.ServiceErrCode, err.Error())
		return
	}
	chatId := chat.ID
	requestId := strings.ReplaceAll(uuid.New().String(), "-", "")
	resp.SetRequestId(requestId)

	message, err := s.repository.SaveMessage(s.ctx, chatId, 0, requestId, req.GetPrompt())
	if err != nil {
		klog.CtxErrorf(s.ctx, "ChatService Run saveMessage err: %+v", err)
		err = errors.NewErrNo(errors.ServiceErrCode, err.Error())
		return
	}

	completion, err := bailian.CreateCompletion(req.GetModelId(), requestId, req.GetSessionId(), req.GetPrompt())
	if err != nil {
		klog.CtxErrorf(s.ctx, "ChatService Run completion err: %+v", err)
		err = errors.NewErrNo(errors.ServiceErrCode, err.Error())
		return
	}
	_, err = s.repository.SaveMessage(s.ctx, chatId, message.Version+1, requestId, completion)
	if err != nil {
		klog.CtxErrorf(s.ctx, "ChatService Run saveMessage err: %+v", err)
		err = errors.NewErrNo(errors.ServiceErrCode, err.Error())
		return
	}
	resp.SetText(completion)
	return
}
