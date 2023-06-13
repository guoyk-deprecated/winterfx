package winterfx

import (
	"go.uber.org/fx"
	"testing"
)

func TestModule(t *testing.T) {
	_ = fx.New(Module)
}
