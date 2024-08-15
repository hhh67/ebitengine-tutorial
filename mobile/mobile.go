package mobile

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"
	"github.com/hhh67/ebitengine-tutorial/game"
)

func init() {
	inogame, err := game.NewGame()
	if err != nil {
		panic(err)
	}
	mobile.SetGame(inogame)
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
