package pacman

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func keyPressed() input {
	// the logic is, we count for which variable the input was more than 0. ebiten the god has written this drivers for keyboard.
	if inpututil.KeyPressDuration(ebiten.KeyUp) > 0 || inpututil.KeyPressDuration(ebiten.KeyK) > 0 {
		return up
	}
	if inpututil.KeyPressDuration(ebiten.KeyLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyH) > 0 {
		return left
	}
	if inpututil.KeyPressDuration(ebiten.KeyRight) > 0 || inpututil.KeyPressDuration(ebiten.KeyL) > 0 {
		return right
	}
	if inpututil.KeyPressDuration(ebiten.KeyDown) > 0 || inpututil.KeyPressDuration(ebiten.KeyJ) > 0 {
		return down
	}
	return 0
}
