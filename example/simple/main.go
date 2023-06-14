package main

import (
	"github.com/guoyk93/winterfx"
	"go.uber.org/fx"
)

type App struct {
}

func CreateApp() *App {
	return &App{}
}

func createRouteHelloWorld() {
}

func main() {
	fx.New(
		winterfx.Module,
		fx.Provide(CreateApp),
	).Run()
}
