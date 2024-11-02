package eblib

var GG = &GameGlobals{}

type GameGlobals struct {
	Game         *Game
	CurrentState Stater
}

func (gg *GameGlobals) SwitchState(newState Stater) {
	gg.Game.SwitchState(newState)
}

func (gg *GameGlobals) GameConfig() *GameConfig {
	return gg.Game.gameConfig
}
