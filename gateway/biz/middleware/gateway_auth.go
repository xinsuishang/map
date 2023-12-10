package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func GatewayAuth() []app.HandlerFunc {
	return []app.HandlerFunc{func(ctx context.Context, c *app.RequestContext) {
		// todo middleware
		return
	}}
}
