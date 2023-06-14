package probefx

import "flag"

type Params struct {
	Cascade int64
}

// DecodeParams decode probe params
func DecodeParams(fset *flag.FlagSet) *Params {
	p := &Params{}
	fset.Int64Var(&p.Cascade, "probe.readiness.cascade", 5, "checker cascade")
	return p
}
