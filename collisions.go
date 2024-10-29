package eblib

import "github.com/solarlune/resolv"

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
func OnOverlap(s1 ISprite, s2 ISprite, overlapAction func(s1 ISprite, s2 ISprite)) {
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
