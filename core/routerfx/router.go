package routerfx

import (
	"github.com/guoyk93/winterfx/core/otelfx"
	"go.uber.org/fx"
	"net/http"
)

type Router interface {
	http.Handler
}

type router struct {
	*Params
	m  *http.ServeMux
	cc chan struct{}
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// concurrency
	if r.cc != nil {
		r.cc <- struct{}{}
		defer func() {
			<-r.cc
		}()
	}

	r.m.ServeHTTP(w, req)
}

type Options struct {
	fx.In

	*Params

	Routes []Route `group:"winterfx_core_routerfx_routes"`
}

func New(opts Options) Router {
	r := &router{
		Params: opts.Params,
		m:      &http.ServeMux{},
	}

	if opts.Concurrency > 0 {
		r.cc = make(chan struct{}, opts.Concurrency)
		for i := 0; i < opts.Concurrency; i++ {
			r.cc <- struct{}{}
		}
	}

	for _, item := range opts.Routes {
		r.m.Handle(item.Pattern,
			otelfx.InstrumentHTTPHandler(
				item.Pattern,
				item.ToHTTPHandler(RouteOptions{
					LoggingRequest:  opts.LoggingRequest,
					LoggingResponse: opts.LoggingResponse,
				}),
			))
	}

	return r
}
