package pacman

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/kgosse/pacmanresources/images"
)

type ghostManager struct {
	ghosts              []*ghost // this is the slice of all the ghosts. we have 8 ghosts each of 4 kinds
	images              map[elem][8]*ebiten.Image
	vulnerabilityImages [5]*ebiten.Image
}

func newGhostManager() *ghostManager {
	gm := &ghostManager{}
	gm.images = make(map[elem][8]*ebiten.Image)
	gm.loadImages()
	return gm
}
func (gm *ghostManager) loadImages() {
	gm.images[blinkyElem] = loadGhostImages(images.BlinkyImages)
	gm.images[clydeElem] = loadGhostImages(images.ClydeImages)
	gm.images[inkyElem] = loadGhostImages(images.InkyImages)
	gm.images[pinkyElem] = loadGhostImages(images.PinkyImages)
	gm.vulnerabilityImages = loadVulenerabilityImages()
}

func (gm *ghostManager) addGhost(x, y int, e elem) {
	gm.ghosts = append(gm.ghosts, newGhost(x, y, e))
}
func (gm *ghostManager) draw(screen *ebiten.Image) {
	// we are iterating through all ghosts - (8 ghosts each of 4 types.)
	for i := 0; i < len(gm.ghosts); i++ {
		g := gm.ghosts[i]
		imgs, _ := gm.images[g.kind]
		//so we are creating map for every kind of ghost here.
		images := make([]*ebiten.Image, 13)
		copy(images, imgs[:]) //[:] -> means entire slice.
		// so each ghost kind will have 8ghosts and 5 vulenerabilities - 4 std vulenerabilities and 1 eaten.png
		copy(images[8:], gm.vulnerabilityImages[:])
		g.draw(screen, images)
	}
}

//  it takes the byte of image and returns the ghost slice.
func loadGhostImages(g [8][]byte) [8]*ebiten.Image {
	var arr [8]*ebiten.Image
	for i := 0; i < 8; i++ {
		img, _, err := image.Decode(bytes.NewReader(g[i]))
		handleError(err)
		arr[i], err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		handleError(err)
	}
	return arr
}
func loadVulenerabilityImages() [5]*ebiten.Image {
	var arr [5]*ebiten.Image
	for i := 0; i < 5; i++ {
		img, _, err := image.Decode(bytes.NewReader(images.VulnerabilityImages[i]))
		handleError(err)
		arr[i], err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		handleError(err)
	}
	return arr
}
func (gm *ghostManager) move(m [][]elem, pac pos) {
	// we will iterate over all ghosts
	for i := 0; i < len(gm.ghosts); i++ {
		g := gm.ghosts[i]
		if g.isMoving() {
			g.findNextMove(m, pac)
		}
		g.move()
	}
}
