package otelfx

import "go.uber.org/fx"

var Module = fx.Module(
	"winterfx_core_otelfx",
	fx.Invoke(Setup),
)
