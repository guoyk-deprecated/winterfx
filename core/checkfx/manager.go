package checkfx

import (
	"context"
	"flag"
	"go.uber.org/fx"
	"sync/atomic"
)

type ManagerParams struct {
	Cascade int64
}

func DecodeManagerParams(fset *flag.FlagSet) *ManagerParams {
	p := &ManagerParams{}
	fset.Int64Var(&p.Cascade, "checker.cascade", 5, "checker cascade")
	return p
}

type Manager interface {
	CheckLiveness() bool

	CheckReadiness(ctx context.Context) (s string, failed bool)
}

type manager struct {
	checkers []Checker

	cascade int64

	failed int64
}

type ManagerOptions struct {
	fx.In

	Params *ManagerParams

	Checkers []Checker `group:"winterfx_core_checkfx"`
}

func AsCheckerBuilder(fn func() Checker) any {
	return fx.Annotate(
		fn,
		fx.ResultTags(`group:"winterfx_core_checkfx"`),
	)
}

func NewManager(opts ManagerOptions) Manager {
	return &manager{
		checkers: opts.Checkers,
		cascade:  opts.Params.Cascade,
	}
}

func (m *manager) CheckLiveness() bool {
	if m.cascade > 0 {
		return m.failed < m.cascade
	} else {
		return true
	}
}

func (m *manager) CheckReadiness(ctx context.Context) (s string, failed bool) {
	r := NewResult()

	for _, c := range m.checkers {
		r.Collect(c.CheckerName(), c.Check(ctx))
	}

	s, failed = r.Result()

	if failed {
		atomic.AddInt64(&m.failed, 1)
	} else {
		atomic.StoreInt64(&m.failed, 0)
	}
	return
}
