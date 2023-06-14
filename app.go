package winterfx

import (
	"github.com/guoyk93/winterfx/core/probefx"
	"github.com/guoyk93/winterfx/core/routerfx"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
	"net/http"
	_ "net/http/pprof"
	"strings"
)

// App the main interface of [summer]
type App interface {
	// Handler inherit [http.Handler]
	http.Handler
}

type app struct {
	*Params

	probe  probefx.Probe
	router routerfx.Router

	hProm http.Handler
}

func (a *app) serveReadiness(rw http.ResponseWriter, req *http.Request) {
	c := routerfx.NewContext(rw, req)
	defer c.Perform()

	s, failed := a.probe.CheckReadiness(c)

	status := http.StatusOK
	if failed {
		status = http.StatusInternalServerError
	}

	c.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	c.Code(status)
	c.Text(s)
}

func (a *app) serveLiveness(rw http.ResponseWriter, req *http.Request) {
	c := routerfx.NewContext(rw, req)
	defer c.Perform()

	c.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	if a.probe.CheckLiveness() {
		c.Code(http.StatusInternalServerError)
		c.Text("CASCADED FAILURE")
	} else {
		c.Code(http.StatusOK)
		c.Text("OK")
	}
}

func (a *app) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// alive, ready, metrics
	if req.URL.Path == a.PathReadiness {
		// support readinessPath == livenessPath
		a.serveReadiness(rw, req)
		return
	} else if req.URL.Path == a.PathLiveness {
		a.serveLiveness(rw, req)
		return
	} else if req.URL.Path == a.PathMetrics {
		a.hProm.ServeHTTP(rw, req)
		return
	}

	// pprof
	if strings.HasPrefix(req.URL.Path, "/debug/pprof") {
		http.DefaultServeMux.ServeHTTP(rw, req)
		return
	}

	// serve with main handler
	a.router.ServeHTTP(rw, req)
}

type Options struct {
	fx.In

	*Params

	probefx.Probe
	routerfx.Router
}

// New create an [App] with [Option]
func New(opts Options) App {
	a := &app{
		Params: opts.Params,
		probe:  opts.Probe,
		router: opts.Router,
	}

	a.hProm = promhttp.Handler()
	return a
}
