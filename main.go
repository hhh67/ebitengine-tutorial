package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hhh67/ebitengine-tutorial/game"
)

func main() {
	game, err := game.NewGame()
	if err != nil {
		panic(err)
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
