package main

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth      = 540
	screenHeight     = 900
	maxParticleCount = 100 // Capped to 100 (for now)
	spriteSize       = 48
	gravity          = 1.2
	acceleration     = 0.5
)

var (
	// Colors

	// Game Variables (Non-constant)
	gameName string = ("")
	game     Game
)

type Game struct {
	GameOver bool
	Dead     bool
	Paused   bool

	Score         int32
	HighScore     int32
	FramesCounter int32

	Particles []Particle
}

type Particle struct {
	Position rl.Vector2
	Color    rl.Color
	Alpha    float32
	Size     float32
	Rotation float32
	Active   bool
}

// NewGame - Start new game
func NewGame() (g Game) {
	g.Init()
	return
}

// On Android this sets callback function to be used for android_main
func init() {
	rl.SetCallbackFunc(main)
}

func main() {
	// Initialize game
	game := NewGame()
	game.GameOver = true

	// Initialize window
	rl.InitWindow(screenWidth, screenHeight, "")

	// Initialize audio
	rl.InitAudioDevice()

	// NOTE: Textures and Sounds MUST be loaded after Window/Audio initialization
	game.Load()

	// Limit FPS
	rl.SetTargetFPS(60)

	// Free resources
	game.Unload()

	// Close audio
	rl.CloseAudioDevice()

	// Close window
	rl.CloseWindow()

	// Exit
	os.Exit(0)
}

// func::Init
// Initializes the game state and all its components.
func (g *Game) Init() {

}

// func::Load
// Loads all resources from a given path.
func (g *Game) Load() {

}

// func::Unload
// Unloads all resources that were previously loaded in.
func (g *Game) Unload() {

}

// func::Update
// Updates the internal game logic and state.
func (g *Game) Update() {

}

// func::Draw
// Draws the game state and all its components to the screen.
func (g *Game) Draw() {

}
