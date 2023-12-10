package notion

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"time"
)

var notionApi = "https://api.notion.com/v1/pages"

type Holder struct {
	header map[string]string
}

func NewNotionHolder(secret string) *Holder {
	return &Holder{
		header: map[string]string{
			consts.HeaderAuthorization: "Bearer " + secret,
			"Notion-Version":           "2022-06-28",
		},
	}
}

func (h *Holder) Do(data any) error {
	jsonData, err := sonic.Marshal(data)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	req := &protocol.Request{}
	res := &protocol.Response{}
	for k, v := range h.header {
		req.Header.Set(k, v)
	}

	req.Header.SetMethod(consts.MethodPost)
	req.Header.SetContentTypeBytes([]byte(consts.MIMEApplicationJSON))
	req.SetRequestURI(notionApi)
	req.SetBody(jsonData)
	c, _ := client.NewClient(
		client.WithMaxConnsPerHost(5),
		client.WithClientReadTimeout(5*time.Second),
		client.WithTLSConfig(&tls.Config{
			MaxVersion: tls.VersionTLS12,
		}),
	)
	err = c.Do(context.Background(), req, res)
	if err != nil {
		return err
	}

	if res.Header.StatusCode() != consts.StatusOK {
		return errors.New(string(res.Body()))
	}

	return nil
}
