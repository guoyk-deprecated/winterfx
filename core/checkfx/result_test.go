package checkfx

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewResult(t *testing.T) {
	r := NewResult()
	r.Collect("redis", errors.New("test"))
	r.Collect("mysql", nil)
	s, failed := r.Result()
	require.True(t, failed)
	require.Equal(t, "redis: test\nmysql: OK", s)
}
