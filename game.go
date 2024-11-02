package eblib

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	state      Stater
	gameConfig *GameConfig
}

type GameConfig struct {
	WindowWidth, WindowHeight, ResolutionWidth, ResolutionHeight int
	WindowTitle                                                  string
}

func DefaultGameConfig(title string) *GameConfig {
	return &GameConfig{
		WindowWidth:      640,
		WindowHeight:     480,
		ResolutionWidth:  160,
		ResolutionHeight: 120,
		WindowTitle:      title,
	}
}

func NewGame(conf *GameConfig, initialState Stater) *Game {
	ebiten.SetWindowSize(conf.WindowWidth, conf.WindowHeight)
	ebiten.SetWindowTitle(conf.WindowTitle)

	g := &Game{
		gameConfig: conf,
	}
	g.state = initialState

	GG.Game = g
	GG.CurrentState = initialState

	return g
}

func (g *Game) Update() error {
	return g.state.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.state.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.gameConfig.ResolutionWidth, g.gameConfig.ResolutionHeight
}

func (g *Game) SwitchState(state Stater) {
	g.state = state
}

func (g *Game) Run() {
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
