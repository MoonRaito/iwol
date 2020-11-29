package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"iWol/ui"
	"log"
)

//go run github.com/hajimehoshi/ebiten/v2/cmd/ebitenmobile bind -target android -javapkg com.hajimehoshi.goinovation -o ./mobile/android/inovation/inovation.aar ./mobile
//go run github.com/hajimehoshi/ebiten/v2/cmd/ebitenmobile bind -target android -javapkg com.moon.iwol -o ./mobile/android/iwol/iwol.aar ./mobile
func main() {
	//game,_ := ui.NewGame()
	ebiten.SetWindowSize(ui.ScreenWidth, ui.ScreenHeight)
	ebiten.SetWindowTitle("UI (Ebiten Demo)")
	if err := ebiten.RunGame(&ui.Uigame); err != nil {
		log.Fatal(err)
	}
}
