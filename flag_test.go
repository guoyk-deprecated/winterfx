package winterfx

import (
	"context"
	"flag"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"os"
	"testing"
)

func TestCreateFlagSet(t *testing.T) {
	fset := CreateFlagSet()
	require.NoError(t, fset.Parse([]string{"--conf", "hello"}))
	f := fset.Lookup("conf")
	require.Equal(t, "hello", f.Value.String())
}

func TestParseFlagSet(t *testing.T) {
	testArgs = []string{"--ignore", "world"}
	fset := CreateFlagSet()
	_ = fset.String("ignore", "", "test")
	val := fset.String("hello", "", "test")
	require.NoError(t, os.Setenv("HELLO", "WORLD"))
	require.NoError(t, ParseFlagSet(ParseFlagSetOptions{FlagSet: fset}))
	require.Equal(t, "WORLD", *val)
}

func TestWrapFlagSetDecoderFunc(t *testing.T) {
	type Params struct {
		Hello string
	}

	testArgs = []string{"--hello", "world"}

	var v string

	app := fx.New(
		fx.Provide(
			CreateFlagSet,
			WrapFlagSetDecoderFunc(func(fset *flag.FlagSet) *Params {
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
