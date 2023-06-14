package checkfx

import (
	"github.com/guoyk93/winterfx/core/flagfx"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"winterfx_core_checkfx",
	fx.Provide(
		flagfx.AsDecoderFunc(DecodeManagerParams),
		NewManager,
	),
)
