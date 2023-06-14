package probefx

import (
	"context"
	"go.uber.org/fx"
)

type CheckerFunc func(ctx context.Context) error

type checker struct {
	name string
	fn   CheckerFunc
}

func AsCheckerProvider[T any](fn func(v T) (name string, cfn CheckerFunc)) any {
	return fx.Annotate(
		func(v T) checker {
			name, cfn := fn(v)
			return checker{name: name, fn: cfn}
		},
		fx.ResultTags(`group:"winterfx_core_probefx_checkers"`),
	)
}
