package pacman

import (
	"bytes"
	"container/list"
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/kgosse/pacmanresources/images"
)

type bigDotManager struct {
	dots   *list.List
	images [2]*ebiten.Image
}

func newBigDotManager() *bigDotManager {
	bd := &bigDotManager{}
	bd.dots = list.New()
	bd.loadImages()
	return bd

}
func (b *bigDotManager) loadImages() {
	imag1, _, err := image.Decode(bytes.NewReader(images.BigDot1_png))
	handleError(err)
	b.images[0], err = ebiten.NewImageFromImage(imag1, ebiten.FilterDefault)
	handleError(err)
	imag2, _, err := image.Decode(bytes.NewReader(images.BigDot2_png))
	handleError(err)
	b.images[1], err = ebiten.NewImageFromImage(imag2, ebiten.FilterDefault)
	handleError(err)
}
func (b *bigDotManager) add(x, y int) {
	b.dots.PushBack(pos{x, y})
}
func (b *bigDotManager) draw(sc *ebiten.Image) {
	for e := b.dots.Front(); e != nil; e = e.Next() {
		d := e.Value.(pos)
		x := float64(d.x * stageBlocSize)
		y := float64(d.y * stageBlocSize)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)
		sc.DrawImage(b.images[0], op)
	}
}
