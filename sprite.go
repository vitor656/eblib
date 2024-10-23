package eblib

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/setanarut/anim"
)

type ISprite interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Sprite struct {
	Img          *ebiten.Image
	X, Y, Dx, Dy float64
	Alive        bool
	Tag          string
	Animations   *anim.AnimationPlayer
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

	s.X += s.Dx
	s.Y += s.Dy
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
