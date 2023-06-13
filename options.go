package winterfx

import (
	"flag"
	"go.uber.org/fx"
	"time"
)

// Params params
type Params struct {
	// Listen listen address
	Listen string

	// Concurrency maximum concurrent requests of [App].
	//
	// A value <= 0 means unlimited
	Concurrency int

	// ReadinessCascade set maximum continuous failed Readiness Checks after which Liveness CheckFunc start to fail.
	ReadinessCascade int64

	// ReadinessPath readiness check path
	ReadinessPath string

	// LivenessPath liveness path
	LivenessPath string

	// MetricsPath metrics path
	MetricsPath string

	// LoggingResponse set loggingResponse
	LoggingResponse bool

	// LoggingRequest set loggingRequest
	LoggingRequest bool

	// DelayStart delay start
	DelayStart time.Duration

	// DelayStop delay stop
	DelayStop time.Duration
}

// ParamsFromFlagSet create Params from flag.FlagSet
func ParamsFromFlagSet(fset *flag.FlagSet) (p *Params) {
	p = &Params{}
	fset.StringVar(&p.Listen, "server.listen", ":8080", "server listen address")
	fset.IntVar(&p.Concurrency, "server.concurrency", 128, "server concurrency")
	fset.Int64Var(&p.ReadinessCascade, "server.readiness.cascade", 5, "server readiness cascade")
	fset.StringVar(&p.ReadinessPath, "server.readiness.path", "/internal/ready", "server path readiness")
	fset.StringVar(&p.LivenessPath, "server.liveness.path", "/internal/alive", "server path liveness")
	fset.StringVar(&p.MetricsPath, "server.metrics.path", "/internal/metrics", "server path metrics")
	fset.BoolVar(&p.LoggingResponse, "server.logging.response", false, "server logging response")
	fset.BoolVar(&p.LoggingRequest, "server.logging.request", false, "server logging request")
	fset.DurationVar(&p.DelayStart, "server.delay.start", time.Second*3, "server delay start")
	fset.DurationVar(&p.DelayStop, "server.delay.stop", time.Second*3, "server delay stop")
	return p
}

type Options struct {
	fx.In
	fx.Lifecycle

	Params
}
