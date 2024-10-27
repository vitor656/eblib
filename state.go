package eblib

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type IState interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type State struct {
	Name           string
	Sprites        []ISprite
	CollisionSpace *resolv.Space
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

	if s.CollisionSpace != nil {
		s.CollisionSpace.Add(sprite.Collider())
	}
}

func (s *State) CreateCollisionSpace(w, h, cw, ch int) {
	s.CollisionSpace = resolv.NewSpace(w, h, cw, ch)
}
