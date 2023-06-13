package main

import (
	"github.com/guoyk93/winterfx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		winterfx.Module,
	).Run()
}
