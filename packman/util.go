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
