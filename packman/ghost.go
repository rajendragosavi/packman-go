package pacman

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

//we have 4 kinds of ghosts - blinky,clyde,pinky,inky.
type ghost struct {
	kind                         elem
	currentImg                   int
	prevPos, currentPos, nextPos pos
	speed                        int
	stepsLength                  pos
	steps                        int
	dir                          input
	vision                       int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func newGhost(x, y int, k elem) *ghost {
	return &ghost{
		kind:        k,
		prevPos:     pos{x, y},
		currentPos:  pos{x, y},
		nextPos:     pos{x, y},
		stepsLength: pos{},
		speed:       2, // ghost's speed is fixed.
		vision:      getVision(k),
	}
}
func getVision(e elem) int {
	switch e {
	case pinkyElem:
		return 10
	case inkyElem:
		return 15
	case blinkyElem:
		return 50
	case clydeElem:
		return 60
	default:
		return 0
	}
}

func (g *ghost) move() {
	switch g.dir {
	case up:
		g.stepsLength.y -= g.speed
	case down:
		g.stepsLength.y += g.speed
	case right:
		g.stepsLength.x += g.speed
	case left:
		g.stepsLength.x -= g.speed
	}
	if g.steps%4 == 0 {
		g.updateImage()
	}
	g.steps++
	if g.steps == 8 {
		g.endMove()
	}
}

// this will return the current ghost image
func (g *ghost) image(imgs []*ebiten.Image) *ebiten.Image {
	return imgs[g.currentImg]
}
func (g *ghost) draw(screen *ebiten.Image, imgs []*ebiten.Image) {
	x := float64(g.currentPos.x*stageBlocSize + g.stepsLength.x)
	y := float64(g.currentPos.y*stageBlocSize + g.stepsLength.y)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(g.image(imgs), op)
}

func (g *ghost) endMove() {
	g.prevPos = g.currentPos
	g.currentPos = g.nextPos
	g.stepsLength = pos{0, 0}
	g.steps = 0
}
func (g *ghost) updateImage() {
	switch g.dir {
	case up:
		if g.currentImg == 6 {
			g.currentImg = 7
		} else {
			g.currentImg = 6
		}
	case down:
		if g.currentImg == 2 {
			g.currentImg = 3
		} else {
			g.currentImg = 2
		}
	case right:
		if g.currentImg == 0 {
			g.currentImg = 1
		} else {
			g.currentImg = 0
		}
	case left:
		if g.currentImg == 4 {
			g.currentImg = 5
		} else {
			g.currentImg = 4
		}
	}
}
func (g *ghost) isMoving() bool {
	if g.steps > 0 {
		return true
	}
	return false

}
func (g *ghost) findNextMove(m [][]elem, pac pos) {
	switch g.localisePlayer(m, pac) {
	case up:
		g.dir = up
	case down:
		g.dir = down
	case right:
		g.dir = right
	case left:
		g.dir = left
	default:
		for _, v := range rand.Perm(5) {
			if v == 0 {
				continue
			}
			dir := input(v)
			np := addPosDir(dir, g.currentPos)
			if canMove(m, np) && np != g.prevPos {
				g.dir = dir
				g.nextPos = np
				return
			}
		}
		g.dir = oppDir(g.dir)
	}
	g.nextPos = addPosDir(g.dir, g.currentPos)
}
func (g *ghost) localisePlayer(m [][]elem, pac pos) input {
	maxY := len(m)
	maxX := len(m[0])

	// up
	if g.currentPos.x == pac.x && g.currentPos.y > pac.y {
		for y, v := g.currentPos.y-1, 1; y >= 0 && v <= g.vision && !isWall(m[y][g.currentPos.x]); y, v = y-1, v+1 {
			if y == pac.y {
				return up
			}
		}
	}

	// down
	if g.currentPos.x == pac.x && g.currentPos.y < pac.y {
		for y, v := g.currentPos.y+1, 1; y < maxY && v <= g.vision && !isWall(m[y][g.currentPos.x]); y, v = y+1, v+1 {
			if y == pac.y {
				return down
			}
		}
	}

	// right
	if g.currentPos.y == pac.y && g.currentPos.x < pac.x {
		for x, v := g.currentPos.x+1, 1; x < maxX && v <= g.vision && !isWall(m[g.currentPos.y][x]); x, v = x+1, v+1 {
			if x == pac.x {
				return right
			}
		}
	}

	// left
	if g.currentPos.y == pac.y && g.currentPos.x > pac.x {
		for x, v := g.currentPos.x-1, 1; x >= 0 && v <= g.vision && !isWall(m[g.currentPos.y][x]); x, v = x-1, v+1 {
			if x == pac.x {
				return left
			}
		}
	}
	return 0
}
