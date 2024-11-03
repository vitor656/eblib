package eblib

var GG = &GameGlobals{}

type GameGlobals struct {
	Game *Game
}

func (gg *GameGlobals) SwitchState(newState Stater) {
	gg.Game.SwitchState(newState)
}

func (gg *GameGlobals) GameConfig() *GameConfig {
	return gg.Game.gameConfig
}

func (gg *GameGlobals) ScreenCenter() (float64, float64) {
	return float64(gg.Game.gameConfig.ResolutionWidth / 2), float64(gg.Game.gameConfig.ResolutionHeight / 2)
}

func (gg *GameGlobals) CurrentState() Stater {
	return gg.Game.state
}

func (gg *GameGlobals) ScreenSize() (int, int) {
	return gg.Game.gameConfig.ResolutionWidth, gg.Game.gameConfig.ResolutionHeight
}
