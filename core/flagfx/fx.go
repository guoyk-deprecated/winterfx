package flagfx

import "go.uber.org/fx"

var Module = fx.Module(
	"winterfx_core_flagfx",
	fx.Provide(
		ArgsFromCommandLine,
		New,
	),
	fx.Invoke(Parse),
)
