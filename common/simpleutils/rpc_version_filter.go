package simpleutils

import (
	"context"
	"github.com/cloudwego/kitex/pkg/discovery"
	"msp/common/constant"
)

func GroupVersionFilter(ctx context.Context, instance []discovery.Instance) []discovery.Instance {
	var res []discovery.Instance

	for _, ins := range instance {
		if match(&ins, constant.Group, getValue(ctx, constant.Group)) && match(&ins, constant.Version, getValue(ctx, constant.Version)) {
			res = append(res, ins)
		}
	}
	return res
}

func match(ins *discovery.Instance, key, value string) bool {
	if v, ok := (*ins).Tag(key); ok {
		return v == value
	}
	return true
}

func getValue(ctx context.Context, value string) (v string) {
	if v, ok := ctx.Value(value).(string); ok {
		return v
	}
	return ""
}
