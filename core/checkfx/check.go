package checkfx

import (
	"context"
)

type Checker interface {
	CheckerName() string

	Check(ctx context.Context) error
}

type CheckerFunc func(ctx context.Context) error

type checker struct {
	n  string
	fn CheckerFunc
}

func (c checker) CheckerName() string {
	return c.n
}

func (c checker) Check(ctx context.Context) error {
	return c.fn(ctx)
}

// NewChecker create a new checker
func NewChecker(name string, fn CheckerFunc) Checker {
	return checker{n: name, fn: fn}
}
