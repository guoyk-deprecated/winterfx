package winterfx

import (
	"flag"
	"time"
)

// Params params
type Params struct {
	// Listen listen address
	Listen string

	// PathReadiness readiness check path
	PathReadiness string

	// PathLiveness liveness path
	PathLiveness string

	// PathMetrics metrics path
	PathMetrics string

	// DelayStart delay start
	DelayStart time.Duration

	// DelayStop delay stop
	DelayStop time.Duration
}

// DecodeParams create Params from flag.FlagSet
func DecodeParams(fset *flag.FlagSet) (p *Params) {
	p = &Params{}
	fset.StringVar(&p.Listen, "server.listen", ":8080", "server listen address")
	fset.StringVar(&p.PathReadiness, "server.path.readiness", "/debug/ready", "server path readiness")
	fset.StringVar(&p.PathLiveness, "server.path.liveness", "/debug/alive", "server path liveness")
	fset.StringVar(&p.PathMetrics, "server.path.metrics", "/debug/metrics", "server path metrics")
	fset.DurationVar(&p.DelayStart, "server.delay.start", time.Second*3, "server delay start")
	fset.DurationVar(&p.DelayStop, "server.delay.stop", time.Second*3, "server delay stop")
	return p
}
