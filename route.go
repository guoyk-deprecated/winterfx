package winterfx

import "go.uber.org/fx"

type Route struct {
	Pattern string
	Action  HandlerFunc
}

func AsRoute(pattern string, fn HandlerFunc) any {
	return fx.Annotate(
		func() Route {
			return Route{
				Pattern: pattern,
				Action:  fn,
			}
		},
		fx.ResultTags(`group:"routes"`),
	)
}
