package winterfx

import (
	"flag"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"testing"
)

func TestModule(t *testing.T) {
	type Params struct {
		Hello string
	}
	app := fx.New(
		Module,
		fx.Decorate(func() FlagSetArgs {
			return FlagSetArgs{"--hello", "world"}
		}),
		fx.Provide(
			WrapFlagSetDecoderFunc(func(fset *flag.FlagSet) *Params {
				p := &Params{}
				fset.StringVar(&p.Hello, "hello", "", "")
				return p
			}),
		),
	)
	require.NoError(t, app.Err())
}
