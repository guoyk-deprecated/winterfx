package probefx

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
	var m Probe

	type testResource struct {
	}

	tr := &testResource{}

	fx.New(
		flagfx.Module,
		Module,

		fx.Provide(
			func() *testResource {
				return tr
			},
			AsCheckerProvider(func(i *testResource) (string, CheckerFunc) {
				return "redis", func(ctx context.Context) error {
					if badRedis {
						return errors.New("test")
					} else {
						return nil
					}
				}
			}),
		),
		fx.Decorate(func() flagfx.Args {
			return []string{
				"--probe.readiness.cascade", "2",
			}
		}),
		fx.Invoke(func(_m Probe) {
			m = _m
		}),
	)

	require.NotNil(t, m)
}
