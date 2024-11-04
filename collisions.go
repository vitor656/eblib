package eblib

import "github.com/solarlune/resolv"

type CollisionManager struct {
	Space *resolv.Space
}

func NewCollisionManager(w, h, cw, ch int) *CollisionManager {
	return &CollisionManager{
		Space: resolv.NewSpace(w, h, cw, ch),
	}
}

func (m *CollisionManager) Add(c Collidable) {
	m.Space.Add(c.Collider())
}

func CollideWithSolid(s *Sprite) {
	if s.Collider() == nil {
		return
	}

	if collision := s.Collider().Check(s.Dx, 0, "solid"); collision != nil {
		s.Dx = collision.ContactWithObject(collision.Objects[0]).X
	}

	if collision := s.Collider().Check(0, s.Dy, "solid"); collision != nil {
		other := collision.Objects[0]
		s.Dy = collision.ContactWithObject(other).Y

		if other.Position.Y > s.Collider().Position.Y {
			s.IsOnGround = true
		}
	} else {
		s.IsOnGround = false
	}

}

// Checks is a sprite overlaps another and trigger funcion once
func OnOverlap(s1 LivableCollidable, s2 LivableCollidable, overlapAction func(s1 any, s2 any)) {
	if s1.Collider() == nil || s2.Collider() == nil {
		return
	}

	if !s1.IsAlive() || !s2.IsAlive() {
		return
	}

	if s1.Collider().Overlaps(s2.Collider()) {
		overlapAction(s1, s2)
	}
}

func PlaceLevelColliders(level *LDTKLevel, layer string, intValue int, space *resolv.Space, w float64, h float64) {
	positions := level.GetIntGridLayerCollisionPoints(layer, intValue)
	for _, p := range positions {
		c := resolv.NewObject(float64(p.X), float64(p.Y), w, h, "solid")
		space.Add(c)
	}
}

func KeepInScreenBounds(s *Sprite) {
	if s.Dx < 0 && s.X < 0 {
		s.Dx = 0
	}

	if s.Dx > 0 && s.X > float64(GG.ScreenWidth()-s.Img.Bounds().Dx()) {
		s.Dx = 0
	}

	if s.Dy < 0 && s.Y < 0 {
		s.Dy = 0
	}

	if s.Dy > 0 && s.Y > float64(GG.ScreenHeight()-s.Img.Bounds().Dy()) {
		s.Dy = 0
	}
}
