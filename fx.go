package winterfx

import (
	"github.com/guoyk93/winterfx/core/checkfx"
	"github.com/guoyk93/winterfx/core/flagfx"
	"go.uber.org/fx"
)

var (
	// Module is the fx module for winterfx
	Module = fx.Module(
		"winterfx",
		flagfx.Module,
		checkfx.Module,
		fx.Provide(
			New,
		),
		fx.Invoke(SetupOTEL),
	)
)
