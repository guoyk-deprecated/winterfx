package routerfx

import (
	"github.com/guoyk93/winterfx/core/flagfx"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"winterfx_core_routefx",
	fx.Provide(
		New,
		flagfx.AsDecoderFunc(DecodeParams),
	),
)
