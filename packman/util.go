package pacman

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
func isWall(e elem) bool {
	if w0 <= e && e <= w24 {
		return true
	}
	return false

}
func loadImage(b []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(b))
	handleError(err)
	ebImg, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	handleError(err)
	return ebImg
}

func canMove(m [][]elem, p pos) bool {
	return !isWall(m[p.y][p.x])
}
func addPosDir(d input, p pos) pos {
	r := pos{p.y, p.x}
	switch d {
	case up:
		r.y--
	case down:
		r.y++
	case right:
		r.x++
	case left:
		r.x--
	}
	if r.x < 0 {
		r.x = 0
	}
	if r.y < 0 {
		r.y = 0
	}
	return r
}
func oppDir(d input) input {
	switch d {
	case up:
		return down
	case left:
		return right
	case down:
		return up
	case right:
		return left
	default:
		return 0
	}
}
