package winterfx

import "go.uber.org/fx"

var (
	// Module is the fx module for winterfx
	Module = fx.Module(
		"winterfx",
		fx.Provide(
			LoadFlagSetArgs,
			NewFlagSet,
			AsFlagSetDecoderFunc(ParamsFromFlagSet),
			New,
		),
		fx.Invoke(ParseFlagSet),
		fx.Invoke(SetupOTEL),
	)
)
