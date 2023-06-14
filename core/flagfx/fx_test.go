package flagfx

import (
	"context"
	"flag"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"testing"
)

func TestModule(t *testing.T) {
	type Params struct {
		Hello string
	}

	var v string

	app := fx.New(
		Module,
		// override args
		fx.Decorate(func() Args {
			return Args{"--hello", "world"}
		}),
		// decode params
		fx.Provide(
			AsDecoderFunc(func(fset *flag.FlagSet) *Params {
				p := &Params{}
				fset.StringVar(&p.Hello, "hello", "", "test")
				return p
			}),
		),
		// extract params
		fx.Invoke(func(p *Params) {
			v = p.Hello
		}),
	)

	require.NoError(t, app.Start(context.Background()))
	require.NoError(t, app.Err())
	require.NoError(t, app.Stop(context.Background()))

	require.Equal(t, "world", v)
}
