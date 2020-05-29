package pacman

import "github.com/hajimehoshi/ebiten"

// Game holds all the pacman game data
type Game struct {
	scene *scene //scene is the game window. we have only one window of game
	in    input
}
type pos struct {
	y, x int
}

// NewGame is a Game constructor
func NewGame() *Game {
	g := &Game{}
	g.scene = newScene(nil)
	return g
}

// ScreenWidth returns the game screen width
func (g *Game) ScreenWidth() int {
	return g.scene.screenWidth()
}

// ScreenHeight returns the game screen height
func (g *Game) ScreenHeight() int {
	return g.scene.screenHeight()
}

// Update updates the screen
func (g *Game) Update(screen *ebiten.Image) error {
	g.in = keyPressed()
	return g.scene.update(screen, g.in)
}
