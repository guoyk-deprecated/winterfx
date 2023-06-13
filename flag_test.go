package winterfx

import (
	"context"
	"flag"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"os"
	"testing"
)

func TestNewFlagSet(t *testing.T) {
	fset := NewFlagSet()
	require.NoError(t, fset.Parse([]string{"--conf", "hello"}))
	f := fset.Lookup("conf")
	require.Equal(t, "hello", f.Value.String())
}

func TestParseFlagSet(t *testing.T) {
	fset := NewFlagSet()
	_ = fset.String("ignore", "", "test")
	val := fset.String("hello", "", "test")
	require.NoError(t, os.Setenv("HELLO", "WORLD"))
	require.NoError(t, ParseFlagSet(ParseFlagSetOptions{FlagSet: fset, Args: FlagSetArgs{"--ignore", "world"}}))
	require.Equal(t, "WORLD", *val)
}

func TestWrapFlagSetDecoderFunc(t *testing.T) {
	type Params struct {
		Hello string
	}

	var v string

	app := fx.New(
		fx.Decorate(func() FlagSetArgs {
			return FlagSetArgs{"--hello", "world"}
		}),
		fx.Provide(
			LoadFlagSetArgs,
			NewFlagSet,
			AsFlagSetDecoderFunc(func(fset *flag.FlagSet) *Params {
				p := &Params{}
				fset.StringVar(&p.Hello, "hello", "", "test")
				return p
			}),
		),
		fx.Invoke(ParseFlagSet),
		fx.Invoke(func(p *Params) {
			v = p.Hello
		}),
	)

	require.NoError(t, app.Start(context.Background()))
	require.NoError(t, app.Err())
	require.NoError(t, app.Stop(context.Background()))

	require.Equal(t, "world", v)
}
