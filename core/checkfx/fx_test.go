package checkfx

import (
	"context"
	"errors"
	"github.com/guoyk93/winterfx/core/flagfx"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"testing"
)

func TestModule(t *testing.T) {
	badRedis := true
	var m Manager

	fx.New(
		flagfx.Module,
		Module,
		fx.Provide(
			AsCheckerBuilder(func() Checker {
				return NewChecker("redis", func(ctx context.Context) error {
					if badRedis {
						return errors.New("test")
					} else {
						return nil
					}
				})
			}),
		),
		fx.Decorate(func() flagfx.Args {
			return []string{
				"--checker.cascade", "2",
			}
		}),
		fx.Invoke(func(_m Manager) {
			m = _m
		}),
	)

	require.NotNil(t, m)
}
