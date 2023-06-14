package probefx

import (
	"context"
	"go.uber.org/fx"
	"sync/atomic"
)

// Probe is a check probe
type Probe interface {
	// CheckLiveness check liveness
	CheckLiveness() bool

	// CheckReadiness check readiness
	CheckReadiness(ctx context.Context) (s string, failed bool)
}

type probe struct {
	*Params

	checkers []checker

	failed int64
}

type Options struct {
	fx.In

	*Params

	Checkers []checker `group:"winterfx_core_probefx_checkers"`
}

func New(opts Options) Probe {
	return &probe{
		checkers: opts.Checkers,
		Params:   opts.Params,
	}
}

func (m *probe) CheckLiveness() bool {
	if m.Cascade > 0 {
		return m.failed < m.Cascade
	} else {
		return true
	}
}

func (m *probe) CheckReadiness(ctx context.Context) (s string, failed bool) {
	r := NewResult()

	for _, c := range m.checkers {
		r.Collect(c.name, c.fn(ctx))
	}

	s, failed = r.Result()

	if failed {
		atomic.AddInt64(&m.failed, 1)
	} else {
		atomic.StoreInt64(&m.failed, 0)
	}
	return
}
