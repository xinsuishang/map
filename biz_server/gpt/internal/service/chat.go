package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"msp/biz_server/gpt/internal/infra/thridpart"
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

	txResource, err := s.repository.WithTx(s.ctx)
	defer func() {
		if v := recover(); v != nil {
			_ = txResource.TxRollback()
			err = errors.NewErrNo(errors.ServiceErrCode, fmt.Errorf("%v", v).Error())
		}
		if err != nil {
			_ = txResource.TxRollback()
		}
	}()
	message, err := txResource.SaveMessage(s.ctx, chatId, 0, requestId, req.GetPrompt())
	if err != nil {
		klog.CtxErrorf(s.ctx, "ChatService Run saveMessage err: %+v", err)
		err = errors.NewErrNo(errors.ServiceErrCode, err.Error())
		return
	}

	completion, err := thridpart.Dispatch(s.ctx, req.GetModelId(), requestId, req.GetSessionId(), req.GetPrompt())
	if err != nil {
		klog.CtxErrorf(s.ctx, "ChatService Run completion err: %+v", err)
		err = errors.NewErrNo(errors.ServiceErrCode, err.Error())
		return
	}
	_, err = txResource.SaveMessage(s.ctx, chatId, message.Version+1, requestId, completion)
	if err != nil {
		klog.CtxErrorf(s.ctx, "ChatService Run saveMessage err: %+v", err)
		err = errors.NewErrNo(errors.ServiceErrCode, err.Error())
		return
	}
	err = txResource.TxCommit()
	if err != nil {
		klog.CtxErrorf(s.ctx, "ChatService Run commit err: %+v", err)
		err = errors.NewErrNo(errors.ServiceErrCode, err.Error())
		return
	}
	resp.SetText(completion)
	return
}
