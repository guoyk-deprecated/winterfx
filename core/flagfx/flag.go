package flagfx

import (
	"flag"
	"github.com/peterbourgon/ff/v3"
	"go.uber.org/fx"
	"os"
)

// Args is the command-line arguments
type Args []string

// ArgsFromCommandLine loads the flag set args from command-line arguments
func ArgsFromCommandLine() Args {
	return os.Args[1:]
}

// JointPoint is a joint point for ensuring all decoder functions are called before parsing flagset
type JointPoint struct{}

type DecoderResult[T any] struct {
	fx.Out
	JointPoint JointPoint `group:"winterfx_core_flagfx_jointpoints"`

	Value T
}

// AsDecoderFunc wraps a flag set decoder function with joint points
func AsDecoderFunc[T any](fn func(fset *flag.FlagSet) T) func(fset *flag.FlagSet) DecoderResult[T] {
	return func(fset *flag.FlagSet) DecoderResult[T] {
		return DecoderResult[T]{
			Value: fn(fset),
		}
	}
}

// New creates a new flag set
func New() *flag.FlagSet {
	name, _ := os.Executable()
	if name == "" {
		name = os.Args[0]
	}
	fset := flag.NewFlagSet(name, flag.ContinueOnError)
	_ = fset.String("conf", "", "config file (optional)")
	return fset
}

// ParseOptions is the options for parsing flag set
type ParseOptions struct {
	fx.In
	JointPoint []JointPoint `group:"winterfx_core_flagfx_jointpoints"`

	FlagSet *flag.FlagSet
	Args    Args
}

// Parse parses the flag set with ff
func Parse(opts ParseOptions) error {
	return ff.Parse(opts.FlagSet, opts.Args,
		ff.WithEnvVars(),
		ff.WithConfigFileFlag("conf"),
		ff.WithConfigFileParser(ff.PlainParser),
	)
}
