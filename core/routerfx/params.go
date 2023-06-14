package routerfx

import "flag"

type Params struct {
	LoggingResponse bool
	LoggingRequest  bool
	Concurrency     int
}

func DecodeParams(fset *flag.FlagSet) *Params {
	p := &Params{}
	fset.BoolVar(&p.LoggingRequest, "router.logging.request", true, "enable request logging")
	fset.BoolVar(&p.LoggingResponse, "router.logging.response", true, "enable response logging")
	fset.IntVar(&p.Concurrency, "router.concurrency", 128, "maximum concurrent requests")
	return p
}
