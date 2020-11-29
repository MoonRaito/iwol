package mobile

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"
	"iWol/ui"
)

func init() {
	//game,_ := ui.NewGame()
	mobile.SetGame(&ui.Uigame)
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
