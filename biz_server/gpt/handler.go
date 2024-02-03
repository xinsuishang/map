package main

import (
	"context"
	gpt "msp/biz_server/kitex_gen/gpt"
)

// ChatServiceImpl implements the last service interface defined in the IDL.
type ChatServiceImpl struct{}

// Chat implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) Chat(ctx context.Context, req *gpt.ApplicationReq) (resp *gpt.ApplicationResp, err error) {
	// TODO: Your code here...
	return
}

// ChatApplicationList implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) ChatApplicationList(ctx context.Context, req *gpt.ApplicationListReq) (resp *gpt.ApplicationListResp, err error) {
	// TODO: Your code here...
	return
}
