package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/vitor656/eblib"
)

func main() {
	state := NewDefaultState()
	g := eblib.NewGame(&eblib.GameConfig{
		WindowWidth:      640,
		WindowHeight:     480,
		ResolutionWidth:  160,
		ResolutionHeight: 120,
		WindowTitle:      "Sprite Example",
	}, state)

	g.Run()
}

type DefaultState struct {
	*eblib.State
	Sprite *eblib.Sprite
}

func NewDefaultState() *DefaultState {
	s := &DefaultState{
		State: eblib.NewState("Default"),
	}

	s.Sprite = eblib.NewSprite(32, 32)
	s.Sprite.MakeSquareImg(16, 16, color.White)
	s.Add(s.Sprite)

	return s
}

func (s *DefaultState) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		s.Sprite.Dx = -2
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		s.Sprite.Dx = 2
	} else {
		s.Sprite.Dx = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeyK) {
		s.Sprite.Kill()
	}

	if ebiten.IsKeyPressed(ebiten.KeyL) {
		s.Sprite.Revive()
	}

	if ebiten.IsKeyPressed(ebiten.Key2) {
		eblib.GG.SwitchState(NewSecondState())
	}

	return s.State.Update()
}

func (s *DefaultState) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Default State")
	s.State.Draw(screen)
}

type SecondState struct {
	*eblib.State
}

func NewSecondState() *SecondState {
	s := &SecondState{
		State: eblib.NewState("Second"),
	}

	return s
}

func (s *SecondState) Update() error {
	if ebiten.IsKeyPressed(ebiten.Key1) {
		eblib.GG.SwitchState(NewDefaultState())
	}

	return s.State.Update()
}

func (s *SecondState) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Second State")
	s.State.Draw(screen)
}
