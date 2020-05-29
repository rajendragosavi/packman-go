package pacman

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/kgosse/pacmanresources/images"
)

type player struct {
	images                    [8]*ebiten.Image
	currentImg                int // index of the current image.
	prevPos, currPos, nextPos pos
	speed                     int
	stepsLength               pos
	steps                     int
	dir                       input
}

func newPlayer(x, y int) *player {
	p := &player{}
	p.loadImages()
	p.currPos = pos{x, y}
	p.prevPos = pos{x, y}
	p.nextPos = pos{x, y}

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
	x := float64(p.currPos.x*stageBlocSize + p.stepsLength.x)
	y := float64(p.currPos.y*stageBlocSize + p.stepsLength.y)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	sc.DrawImage(p.image(), op) // draw the image with current image of the player.
}

func (p *player) move(m [][]elem, dir input) {
	// not moving and no direction
	if !p.isMoving() && dir == 0 {
		return
	}
	//new direction,
	// if we are not moving but we got a direction then..
	if !p.isMoving() && dir != 0 {
		if !canMove(m, addPosDir(dir, p.currPos)) {
			return
		}
		p.updateDir(dir)
	}
	// lets adjust the speed. we have divided the speed part
	if p.steps <= 1 || p.steps >= 6 {
		p.speed = 4
	} else {
		p.speed = 5
	}
	// move and update the direction.
	switch p.dir {
	case up:
		p.stepsLength.y -= p.speed
	case right:
		p.stepsLength.x += p.speed
	case left:
		p.stepsLength.x -= p.speed
	case down:
		p.stepsLength.y -= p.speed
	}
	if p.steps > 5 {
		p.updateImage(false)
	} else {
		p.updateImage(true)
	}
	p.steps++
	if p.steps >= 7 {
		p.endMove()
	}

}

func (p *player) isMoving() bool {
	if p.steps > 0 {
		return true
	}
	return false
}
func (p *player) updateDir(d input) {
	p.stepsLength = pos{0, 0}
	p.dir = d
	p.nextPos = addPosDir(d, p.currPos)
	p.prevPos = p.currPos
}

func (p *player) updateImage(openMouth bool) {
	switch p.dir {
	case up:
		if openMouth {
			p.currentImg = 7
		} else {
			p.currentImg = 6
		}
	case right:
		if openMouth {
			p.currentImg = 1
		} else {
			p.currentImg = 0
		}
	case down:
		if openMouth {
			p.currentImg = 3
		} else {
			p.currentImg = 2
		}
	case left:
		if openMouth {
			p.currentImg = 5
		} else {
			p.currentImg = 4
		}
	}
}
func (p *player) endMove() {
	p.currPos = p.nextPos
	p.stepsLength = pos{0, 0}
	p.steps = 0
}
