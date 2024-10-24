package eblib

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/setanarut/anim"
	"github.com/solarlune/resolv"
)

type ISprite interface {
	Update()
	Draw(screen *ebiten.Image)
	GetCollider() *resolv.Object
}

type Sprite struct {
	Img          *ebiten.Image
	X, Y, Dx, Dy float64
	Alive        bool
	Tag          string
	Animations   *anim.AnimationPlayer
	Collider     *resolv.Object
}

func NewSprite(x, y float64) *Sprite {
	s := &Sprite{
		X:     x,
		Y:     y,
		Alive: true,
	}

	return s
}

func (s *Sprite) Update() {
	if !s.Alive {
		return
	}

	if s.Collider != nil {
		s.checkCollisionSolid()
		s.X += s.Dx
		s.Y += s.Dy
		s.Collider.Position.X = s.X
		s.Collider.Position.Y = s.Y
		s.Collider.Update()
	} else {
		s.X += s.Dx
		s.Y += s.Dy
	}

}

func (s *Sprite) Draw(screen *ebiten.Image) {
	if !s.Alive {
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
}

func (s *Sprite) Revive() {
	s.Alive = true
}

func (s *Sprite) SetupAnimatedSprite(spritesheet *ebiten.Image) {
	s.Animations = anim.NewAnimationPlayer(spritesheet)
}

func (s *Sprite) SetCollider(x, y, w, h float64, tags ...string) {
	s.Collider = resolv.NewObject(x, y, w, h, tags...)
}

func (s *Sprite) GetCollider() *resolv.Object {
	return s.Collider
}

func (s *Sprite) checkCollisionSolid() {
	if s.Collider == nil {
		return
	}

	if collision := s.Collider.Check(s.Dx, 0, "solid"); collision != nil {
		s.Dx = collision.ContactWithObject(collision.Objects[0]).X
	}

	if collision := s.Collider.Check(0, s.Dy, "solid"); collision != nil {
		s.Dy = collision.ContactWithObject(collision.Objects[0]).Y
	}
}
