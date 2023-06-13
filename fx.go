package winterfx

import "go.uber.org/fx"

var (
	// Module is the fx module for winterfx
	Module = fx.Module(
		"winterfx",
		fx.Provide(
			CreateFlagSet,
		),
		fx.Invoke(ParseFlagSet),
	)
)
