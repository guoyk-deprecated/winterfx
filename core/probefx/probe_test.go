package probefx

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewManager(t *testing.T) {
	badRedis := true
	m := New(Options{
		Params: &Params{
			Cascade: 2,
		},
		Checkers: []checker{
			{
				name: "redis",
				fn: func(ctx context.Context) error {
					if badRedis {
						return errors.New("test")
					}
					return nil
				},
			},
			{
				name: "mysql",
				fn: func(ctx context.Context) error {
					return nil
				},
			},
		},
	})
	require.True(t, m.CheckLiveness())

	s, failed := m.CheckReadiness(context.Background())
	require.True(t, failed)
	require.Equal(t, "redis: test\nmysql: OK", s)
	require.True(t, m.CheckLiveness())

	s, failed = m.CheckReadiness(context.Background())
	require.True(t, failed)
	require.Equal(t, "redis: test\nmysql: OK", s)
	require.False(t, m.CheckLiveness())

	badRedis = false

	s, failed = m.CheckReadiness(context.Background())
	require.False(t, failed)
	require.Equal(t, "redis: OK\nmysql: OK", s)
	require.True(t, m.CheckLiveness())

}
