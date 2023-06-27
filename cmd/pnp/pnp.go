package main

import (
	_ "embed"

	"github.com/ronna-s/go-ood/pkg/pnpdev"

	"github.com/ronna-s/go-ood/pkg/pnp"
	engine "github.com/ronna-s/go-ood/pkg/pnp/engine/tview"
)

func main() {
	game := pnp.New(
		pnpdev.NewZombieGopher(),
		pnpdev.NewMinion(),
		pnpdev.NewRubyist(),
		pnpdev.NewGopher(),
	)
	game.Run(engine.New())
}
