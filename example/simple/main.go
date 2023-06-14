package main

import (
	"context"
	"github.com/guoyk93/winterfx"
	"github.com/guoyk93/winterfx/core/probefx"
	"github.com/guoyk93/winterfx/core/routerfx"
	"go.uber.org/fx"
)

type App struct {
}

func CreateApp() *App {
	return &App{}
}

func createCheckTest(a *App) (string, probefx.CheckerFunc) {
	return "test", func(ctx context.Context) error {
		return nil
	}
}

func createRouteHelloWorld(a *App) (string, routerfx.HandlerFunc) {
	return "/hello", func(c routerfx.Context) {
		c.Text("world")
	}
}

func main() {
	fx.New(
		winterfx.Module,
		fx.Provide(
			CreateApp,
			routerfx.AsRouteProvider(createRouteHelloWorld),
			probefx.AsCheckerProvider(createCheckTest),
		),
	).Run()
}
