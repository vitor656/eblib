package eblib

import (
	"image/color"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/setanarut/anim"
	"github.com/solarlune/resolv"
)

type Sprite struct {
	State           Stater
	Img             *ebiten.Image
	X, Y, Dx, Dy    float64
	Gravity         float64
	Alive           bool
	Tag             string
	Animations      *anim.AnimationPlayer
	collider        *resolv.Object
	ColliderOffsetX float64
	ColliderOffsetY float64
	IsOnGround      bool
	id              uuid.UUID
}

func NewSprite(x, y float64) *Sprite {
	s := &Sprite{
		X:     x,
		Y:     y,
		Alive: true,
		id:    uuid.New(),
	}

	return s
}

func (s *Sprite) ID() uuid.UUID {
	return s.id
}

func (s *Sprite) Update() error {
	if !s.Alive {
		return nil
	}

	s.Dy += s.Gravity

	s.X += s.Dx
	s.Y += s.Dy

	s.updateCollider()

	return nil
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	if !s.Alive {
		return
	}

	if !s.IsOnScreen() {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.X, s.Y)
	screen.DrawImage(s.Img, op)
}

func (s *Sprite) MakeSquareImg(w, h int, c color.Color) {
	s.Img = ebiten.NewImage(w, h)
	s.Img.Fill(c)
}

func (s *Sprite) Kill() {
	s.Alive = false

	space := s.collider.Space
	space.Remove(s.collider)

	state := s.State
	state.Remove(s)
}

func (s *Sprite) Revive() {
	s.Alive = true
}

func (s *Sprite) IsAlive() bool {
	return s.Alive
}

func (s *Sprite) SetupAnimatedSprite(spritesheet *ebiten.Image) {
	s.Animations = anim.NewAnimationPlayer(spritesheet)
}

func (s *Sprite) SetCollider(x, y, w, h float64, tags ...string) {
	s.collider = resolv.NewObject(x, y, w, h, tags...)
}

func (s *Sprite) SetColliderData(data interface{}) {
	if s.collider == nil {
		return
	}

	s.collider.Data = data
}

func (s *Sprite) Collider() *resolv.Object {
	return s.collider
}

func (s *Sprite) updateCollider() {
	if s.collider == nil {
		return
	}

	s.collider.Position.X = s.X + s.ColliderOffsetX
	s.collider.Position.Y = s.Y + s.ColliderOffsetY
	s.collider.Update()
}

func (s *Sprite) ScreenCenter() {
	s.X, s.Y = GG.ScreenCenter()

	s.X = s.X - float64(s.Img.Bounds().Dx())/2
	s.Y = s.Y - float64(s.Img.Bounds().Dy())/2
}

func (s *Sprite) Size() (int, int) {
	p := s.Img.Bounds().Size()
	return p.X, p.Y
}

func (s *Sprite) SetState(state Stater) {
	s.State = state
}

func (s *Sprite) IsOnScreen() bool {
	screenW, screenH := GG.ScreenSize()
	w, h := s.Size()
	if s.X+float64(w) < 0 || s.X > float64(screenW) || s.Y+float64(h) < 0 || s.Y > float64(screenH) {
		return false
	}
	return true
}
