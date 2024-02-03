package errors

import (
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"testing"
)

func TestBizErrorStatus(t *testing.T) {
	err, ok := kerrors.FromBizStatusError(Success)

	assert.Assert(t, ok)
	assert.Assert(t, errors.Is(err, Success))
	err2 := fmt.Errorf("ent: %s", "test")
	bizErr, ok := kerrors.FromBizStatusError(err2)
	assert.Assert(t, !ok)
	assert.Assert(t, errors.Is(nil, bizErr))
}
