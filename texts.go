package eblib

import (
	"bytes"
	"image/color"
	"io"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Text struct {
	Text        string
	X, Y        float64
	Color       color.Color
	Size        float64
	textSource  *text.GoTextFaceSource
	DrawOptions *text.DrawOptions
}

func NewText(content string) *Text {
	t := &Text{
		Text:  content,
		Size:  8,
		Color: color.White,
	}

	// use default lib font
	t.SetupFont(bytes.NewReader(MonoFont3x3))

	// t.textSource = source
	t.DrawOptions = &text.DrawOptions{}

	return t
}

func (t *Text) Draw(screen *ebiten.Image) {
	t.DrawOptions.GeoM.Reset()
	t.DrawOptions.ColorScale.Reset()
	t.DrawOptions.GeoM.Translate(t.X, t.Y)
	t.DrawOptions.ColorScale.ScaleWithColor(t.Color)
	text.Draw(screen, t.Text, &text.GoTextFace{Source: t.textSource, Size: t.Size}, t.DrawOptions)
}

// Use SetupFont(bytes.NewReader(embedFont))
func (t *Text) SetupFont(source io.Reader) {
	s, err := text.NewGoTextFaceSource(source)
	if err != nil {
		panic(err)
	}

	t.textSource = s
}
