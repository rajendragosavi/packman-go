package pacman

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/kgosse/pacmanresources/images"
)

type player struct {
	images     [8]*ebiten.Image
	currentImg int // index of the current image.
	currPos    pos
}

func newPlayer(x, y int) *player {
	p := &player{}
	p.loadImages()
	p.currPos = pos{x, y}
	return p
}

// we have 8 possible images for player position
func (p *player) loadImages() {
	for i := 0; i < 8; i++ {
		img, _, err := image.Decode(bytes.NewReader(images.PlayerImages[i]))
		handleError(err)
		p.images[i], err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		handleError(err)
	}
}
func (p *player) image() *ebiten.Image {
	return p.images[p.currentImg] // it returns the current image of player whenever invoked . it is used to get the player image at any given time.
}

// this is to draw the image of the player .
func (p *player) draw(sc *ebiten.Image) {
	x := float64(p.currPos.x * stageBlocSize)
	y := float64(p.currPos.y * stageBlocSize)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	sc.DrawImage(p.image(), op) // draw the image with current image of the player.
}
