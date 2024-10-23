package eblib

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type IState interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type State struct {
	Name    string
	Sprites []ISprite
}

func NewState(name string) *State {
	s := &State{}
	s.Name = name
	s.Sprites = make([]ISprite, 0)

	return s
}

func (s *State) Update() error {
	for _, sprite := range s.Sprites {
		sprite.Update()
	}
	return nil
}

func (s *State) Draw(screen *ebiten.Image) {
	for _, sprite := range s.Sprites {
		sprite.Draw(screen)
	}
}

func (s *State) Add(sprite ISprite) {
	s.Sprites = append(s.Sprites, sprite)
}
