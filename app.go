package winterfx

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"net/http"
	"net/http/pprof"
	"strings"
	"sync/atomic"
)

// HandlerFunc handler func with [Context] as argument
type HandlerFunc func(c Context)

// App the main interface of [summer]
type App interface {
	// Handler inherit [http.Handler]
	http.Handler

	// HandleFunc register an action function with given path pattern
	//
	// This function is similar with [http.ServeMux.HandleFunc]
	HandleFunc(pattern string, fn HandlerFunc)
}

type app struct {
	Params

	hMain *http.ServeMux

	hProm http.Handler
	hProf http.Handler

	cc chan struct{}

	failed int64
}

func (a *app) HandleFunc(pattern string, fn HandlerFunc) {
	a.hMain.Handle(
		pattern,
		otelhttp.NewHandler(
			otelhttp.WithRouteTag(
				pattern,
				http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					c := newContext(rw, req)
					c.loggingResponse = a.LoggingResponse
					func() {
						defer c.Perform()
						fn(c)
					}()
				}),
			),
			pattern,
		),
	)
}

func (a *app) serveReadiness(rw http.ResponseWriter, req *http.Request) {
	c := newContext(rw, req)
	defer c.Perform()

	cr := NewCheckResult()
	//TODO: add readiness check
	s, failed := cr.Result()

	status := http.StatusOK
	if failed {
		atomic.AddInt64(&a.failed, 1)
		status = http.StatusInternalServerError
	} else {
		atomic.StoreInt64(&a.failed, 0)
	}

	c.Code(status)
	c.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Text(s)
}

func (a *app) serveLiveness(rw http.ResponseWriter, req *http.Request) {
	c := newContext(rw, req)
	defer c.Perform()

	c.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	if a.ReadinessCascade > 0 &&
		atomic.LoadInt64(&a.failed) > a.ReadinessCascade {
		c.Code(http.StatusInternalServerError)
		c.Text("CASCADED")
	} else {
		c.Code(http.StatusOK)
		c.Text("OK")
	}
}

func (a *app) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// alive, ready, metrics
	if req.URL.Path == a.ReadinessPath {
		// support readinessPath == livenessPath
		a.serveReadiness(rw, req)
		return
	} else if req.URL.Path == a.LivenessPath {
		a.serveLiveness(rw, req)
		return
	} else if req.URL.Path == a.MetricsPath {
		a.hProm.ServeHTTP(rw, req)
		return
	}

	// pprof
	if strings.HasPrefix(req.URL.Path, "/debug/pprof") {
		a.hProf.ServeHTTP(rw, req)
		return
	}

	// concurrency
	if a.cc != nil {
		<-a.cc
		defer func() {
			a.cc <- struct{}{}
		}()
	}

	// serve with main handler
	a.hMain.ServeHTTP(rw, req)
}

// New create an [App] with [Option]
func New(opts Options) App {
	a := &app{
		Params: opts.Params,
	}

	// create handlers
	{
		a.hMain = &http.ServeMux{}
		a.hProm = promhttp.Handler()
		m := &http.ServeMux{}
		m.HandleFunc("/debug/pprof/", pprof.Index)
		m.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		m.HandleFunc("/debug/pprof/profile", pprof.Profile)
		m.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		m.HandleFunc("/debug/pprof/trace", pprof.Trace)
		a.hProf = m
	}

	// create concurrency controller
	if a.Concurrency > 0 {
		a.cc = make(chan struct{}, a.Concurrency)
		for i := 0; i < a.Concurrency; i++ {
			a.cc <- struct{}{}
		}
	}
	return a
}
