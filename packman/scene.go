package pacman

import (
	"github.com/hajimehoshi/ebiten"

	pacimages "github.com/kgosse/pacmanresources/images"
)

type scene struct {
	matrix        [][]elem
	wallSurface   *ebiten.Image
	images        map[elem]*ebiten.Image
	stage         *stage
	dotManager    *dotManager
	bigDotManager *bigDotManager
	player        *player
	ghostManager  *ghostManager
	textManager   *textManager
}

func newScene(st *stage) *scene {
	s := scene{}
	s.stage = st
	if s.stage == nil {
		s.stage = defaultStage
	}
	s.images = make(map[elem]*ebiten.Image)
	s.dotManager = newDotManager()
	s.bigDotManager = newBigDotManager()
	s.ghostManager = newGhostManager()
	h := len(s.stage.matrix)
	w := len(s.stage.matrix[0])
	s.textManager = newTextManager(w*stageBlocSize, h*stageBlocSize)
	s.loadImages()       // initializes the images
	s.createStage()      // initialize the matrix
	s.buildWallSurface() // intialize the surface.
	return &s
}
func (s *scene) update(screen *ebiten.Image, in input) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	s.player.move(s.matrix, in)
	s.ghostManager.move(s.matrix, s.player.currPos)
	screen.Clear()
	screen.DrawImage(s.wallSurface, nil)
	s.dotManager.draw(screen)
	s.bigDotManager.draw(screen)
	s.player.draw(screen)
	s.ghostManager.draw(screen)

	s.textManager.draw(screen, 0, 1, s.player.images[1]) //score ,lives and players we are hardcoding here
	return nil
}

// this sets the stage for the game.
func (s *scene) createStage() {
	h := len(s.stage.matrix)
	w := len(s.stage.matrix[0])
	s.matrix = make([][]elem, h)
	for i := 0; i < h; i++ {
		s.matrix[i] = make([]elem, w)
		for j := 0; j < w; j++ {
			c := s.stage.matrix[i][j] - '0'
			if c <= 9 {
				s.matrix[i][j] = elem(c)
			} else {
				s.matrix[i][j] = elem(s.stage.matrix[i][j] - 'a' + 10)
			}
			// we are going to put dots on the stage.
			// covention , smalldot -r , bigdot - q , player - p check the stage.go file for reference.

			switch s.matrix[i][j] {
			case dotElem: // check if it dot cordinate and give it to add funciton, which will add a dot on stage
				s.dotManager.add(i, j)
			case bigDotElem: // same logic as dot.
				s.bigDotManager.add(i, j)
			case playerElem:
				s.player = newPlayer(i, j)
			case blinkyElem: //  whenever we  came across  a ghost , we add that into ghostslice.
				s.ghostManager.addGhost(i, j, blinkyElem)
			case inkyElem:
				s.ghostManager.addGhost(i, j, inkyElem)
			case clydeElem:
				s.ghostManager.addGhost(i, j, clydeElem)
			case pinkyElem:
				s.ghostManager.addGhost(i, j, pinkyElem)
			}
		}
	}
}

// this will return screen width
func (s *scene) screenWidth() int {
	w := len(s.stage.matrix[0]) // row of the stage
	return w * stageBlocSize    //stageBlocksize is 32 (|=============width=============|)
}

// this will return screen height
func (s *scene) screenHeight() int {
	h := len(s.stage.matrix)
	sizeH := ((h*stageBlocSize)/backgroundImageSize + 2) * backgroundImageSize
	return sizeH // same it will return width
}

func (s *scene) loadImages() {
	for i := w0; i <= w24; i++ {
		s.images[i] = loadImage(pacimages.WallImages[i])
	}
	s.images[backgroundElem] = loadImage(pacimages.Background_png)
}

// need to dig deep
func (s *scene) buildWallSurface() {
	h := len(s.stage.matrix)
	w := len(s.stage.matrix[0])
	sizeW := ((w*stageBlocSize)/backgroundImageSize + 1) * backgroundImageSize
	sizeH := ((h*backgroundImageSize)/backgroundImageSize + 2) * backgroundImageSize // increase the height for text

	s.wallSurface, _ = ebiten.NewImage(sizeW, sizeH, ebiten.FilterDefault)
	for i := 0; i < sizeH/backgroundImageSize; i++ {
		y := float64(i * backgroundImageSize)
		for j := 0; j < sizeW/backgroundImageSize; j++ {
			op := &ebiten.DrawImageOptions{}
			x := float64(j * backgroundImageSize)
			op.GeoM.Translate(x, y)
			err := s.wallSurface.DrawImage(s.images[backgroundElem], op)
			handleError(err)
		}
	}
	for i := 0; i < h; i++ {
		y := float64(i * stageBlocSize)
		for j := 0; j < w; j++ {
			//passing every element to check if it is wall or not.
			if !isWall(s.matrix[i][j]) {
				continue
			}
			op := &ebiten.DrawImageOptions{}
			x := float64(j * stageBlocSize)
			op.GeoM.Translate(x, y)
			err := s.wallSurface.DrawImage(s.images[s.matrix[i][j]], op)
			handleError(err)
		}
	}

}
