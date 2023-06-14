package routerfx

import (
	"github.com/guoyk93/winterfx/core/flagfx"
	"go.uber.org/fx"
	"testing"
)

func TestModule(t *testing.T) {
	type res struct {
	}

	r := &res{}
	var router Router

	app := fx.New(
		flagfx.Module,
		Module,
		flagfx.OverrideArgs([]string{"-router.concurrency=2"}),
		fx.Supply(r),
		fx.Provide(
			AsRouteProvider(func(r *res) (string, HandlerFunc) {
				return "/hello", func(c Context) {
					c.Text("world")
				}
			}),
		),
		fx.Populate(&router),
	)
	_ = app
}
