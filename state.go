package eblib

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type State struct {
	Name       string
	Sprites    []UpdatableDrawable
	colManager *CollisionManager
}

func NewState(name string) *State {
	s := &State{}
	s.Name = name
	s.Sprites = make([]UpdatableDrawable, 0)

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

func (s *State) Add(sprite UpdatableDrawable) {
	s.Sprites = append(s.Sprites, sprite)

	if collidable, ok := sprite.(Collidable); ok && s.colManager != nil {
		s.colManager.Add(collidable)
	}
}

func (s *State) CreateCollisionSpace(w, h, cw, ch int) {
	s.colManager = NewCollisionManager(w, h, cw, ch)
}

func (s *State) CollisionManager() *CollisionManager {
	return s.colManager
}
