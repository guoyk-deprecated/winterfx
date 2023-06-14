package routerfx

import (
	"go.uber.org/fx"
	"net/http"
)

// HandlerFunc handler func with [Context] as argument
type HandlerFunc func(c Context)

// Route a route with pattern and handler
type Route struct {
	Pattern string
	HandlerFunc
}

type RouteOptions struct {
	LoggingRequest  bool
	LoggingResponse bool
}

func (r Route) ToHTTPHandler(opts RouteOptions) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		c := newContext(rw, req)
		c.loggingResponse = opts.LoggingResponse
		c.loggingRequest = opts.LoggingRequest
		func() {
			defer c.Perform()
			r.HandlerFunc(c)
		}()
	})
}

func AsRouteProvider[T any](fn func(v T) (pattern string, h HandlerFunc)) any {
	return fx.Annotate(
		func(v T) Route {
			pattern, rfn := fn(v)
			return Route{Pattern: pattern, HandlerFunc: rfn}
		},
		fx.ResultTags(`group:"winterfx_core_routerfx_routes"`),
	)
}
