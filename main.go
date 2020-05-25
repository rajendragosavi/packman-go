package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	pacman "github.com/rajendragosavi/packman-go/packman"
)

func main() {

	g := pacman.NewGame() //return game instance.
	if err := ebiten.Run(g.Update, g.ScreenWidth(), g.ScreenHeight(), 1, "Pacman"); err != nil {

		log.Fatal(err)
	}

}
