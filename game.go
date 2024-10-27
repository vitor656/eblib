package eblib

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	State      IState
	GameConfig *GameConfig
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

func NewGame(conf *GameConfig, initialState IState) *Game {
	ebiten.SetWindowSize(conf.WindowWidth, conf.WindowHeight)
	ebiten.SetWindowTitle(conf.WindowTitle)

	g := &Game{
		GameConfig: conf,
	}
	g.State = initialState

	GameInstance = g
	return g
}

func (g *Game) Update() error {
	return g.State.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.State.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.GameConfig.ResolutionWidth, g.GameConfig.ResolutionHeight
}

func (g *Game) SwitchState(state IState) {
	g.State = state
}

func (g *Game) Run() {
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
