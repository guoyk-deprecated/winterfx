package winterfx

import (
	"github.com/guoyk93/winterfx/core/flagfx"
	"github.com/guoyk93/winterfx/core/otelfx"
	"github.com/guoyk93/winterfx/core/probefx"
	"go.uber.org/fx"
)

var (
	// Module is the fx module for winterfx
	Module = fx.Module(
		"winterfx",
		flagfx.Module,
		probefx.Module,
		otelfx.Module,
		fx.Provide(
			flagfx.AsDecoderFunc(DecodeParams),
			New,
		),
	)
)
