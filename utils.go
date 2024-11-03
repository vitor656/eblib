package eblib

import (
	_ "embed"
	"image/color"
)

//go:embed fonts/3X3Mono-drx1V.ttf
var MonoFont3x3 []byte

//go:embed fonts/PressStart2P-vaV7.ttf
var PressStart2P []byte

var (
	Color3310A_Lighter = color.RGBA{R: 78, G: 94, B: 85, A: 255}
	Color3310A_Darker  = color.RGBA{R: 26, G: 32, B: 24, A: 255}
	Color3310B_Lighter = color.RGBA{R: 61, G: 78, B: 0, A: 255}
	Color3310B_Darker  = color.RGBA{R: 17, G: 25, B: 4, A: 255}
)
