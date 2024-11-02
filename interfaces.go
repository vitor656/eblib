package eblib

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type Updatable interface {
	Update() error
}

type Drawable interface {
	Draw(screen *ebiten.Image)
}

type Collidable interface {
	Collider() *resolv.Object
}

type Livable interface {
	IsAlive() bool
}

type LivableCollidable interface {
	Livable
	Collidable
}

type UpdatableDrawable interface {
	Updatable
	Drawable
}

type Stater interface {
	Updatable
	Drawable
}
