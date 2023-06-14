package probefx

import (
	"github.com/guoyk93/winterfx/core/flagfx"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"winterfx_core_probefx",
	fx.Provide(
		flagfx.AsDecoderFunc(DecodeParams),
		New,
	),
)
