package pacman

import "github.com/hajimehoshi/ebiten"

//we have 4 kinds of ghosts - blinky,clyde,pinky,inky.
type ghost struct {
	kind       elem
	currentImg int
	currentPos pos
}

func newGhost(x, y int, k elem) *ghost {
	return &ghost{
		kind:       k,
		currentPos: pos{x, y},
	}
}

// this will return the current ghost image
func (g *ghost) image(imgs []*ebiten.Image) *ebiten.Image {
	return imgs[g.currentImg]
}
func (g *ghost) draw(screen *ebiten.Image, imgs []*ebiten.Image) {
	x := float64(g.currentPos.x * stageBlocSize)
	y := float64(g.currentPos.y * stageBlocSize)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(g.image(imgs), op)
}
