package eblib

var GameInstance *Game

func SwitchState(newState IState) {
	GameInstance.SwitchState(newState)
}
