package winterfx

import (
	"flag"
	"github.com/peterbourgon/ff/v3"
	"go.uber.org/fx"
	"os"
)

// FlagSetJointPoint is a joint point for ensuring all decoder functions are called before parsing flagset
type FlagSetJointPoint struct{}

// CreateFlagSet creates a new flag set
func CreateFlagSet() *flag.FlagSet {
	name, _ := os.Executable()
	if name == "" {
		name = os.Args[0]
	}
	fset := flag.NewFlagSet(name, flag.ContinueOnError)
	_ = fset.String("conf", "", "config file (optional)")
	return fset
}

type FlagSetDecodeResult[T any] struct {
	fx.Out
	FlagSetJointPoint `group:"flagsetjointpoint"`

	Value T
}

// WrapFlagSetDecoderFunc wraps a flag set decoder function with joint points
func WrapFlagSetDecoderFunc[T any](fn func(fset *flag.FlagSet) T) func(fset *flag.FlagSet) FlagSetDecodeResult[T] {
	return func(fset *flag.FlagSet) FlagSetDecodeResult[T] {
		return FlagSetDecodeResult[T]{
			FlagSetJointPoint: FlagSetJointPoint{},
			Value:             fn(fset),
		}
	}
}

type FlagSetArgs []string

// LoadFlagSetArgs loads the flag set args
func LoadFlagSetArgs() FlagSetArgs {
	return os.Args[1:]
}

// ParseFlagSetOptions is the options for parsing flag set
type ParseFlagSetOptions struct {
	fx.In
	JointPoints []FlagSetJointPoint `group:"flagsetjointpoint"`

	FlagSet *flag.FlagSet

	Args FlagSetArgs
}

// ParseFlagSet parses the flag set
func ParseFlagSet(opts ParseFlagSetOptions) error {
	return ff.Parse(opts.FlagSet, opts.Args,
		ff.WithEnvVars(),
		ff.WithConfigFileFlag("conf"),
		ff.WithConfigFileParser(ff.PlainParser),
	)
}
