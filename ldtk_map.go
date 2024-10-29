package eblib

import (
	"image"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/ldtkgo"
	renderer "github.com/solarlune/ldtkgo/renderer/ebitengine"
)

type LDTKLevel struct {
	lDTKProject  *ldtkgo.Project
	renderer     *renderer.Renderer
	currentLevel int
}

// Requires a go:embed assets folder
func NewLDTKLevel(projectPath string, assets fs.FS) *LDTKLevel {
	level := &LDTKLevel{}

	proj, err := ldtkgo.Open(projectPath, assets)
	if err != nil {
		panic(err)
	}

	level.lDTKProject = proj

	subDir, err := fs.Sub(assets, "assets")

	if err != nil {
		panic(err)
	}

	level.renderer, err = renderer.New(subDir, level.lDTKProject)

	if err != nil {
		panic(err)
	}

	return level
}

func (m *LDTKLevel) SwitchLevel(level int) {
	m.currentLevel = level
}

func (m *LDTKLevel) Draw(screen *ebiten.Image) {
	level := m.lDTKProject.Levels[m.currentLevel]
	opt := renderer.NewDefaultDrawOptions()
	m.renderer.Render(level, screen, opt)
}

func (m *LDTKLevel) GetIntGridLayerCollisionPoints(layerName string, value int) []image.Point {
	data := make([]image.Point, 0)
	layer := m.lDTKProject.Levels[m.currentLevel].LayerByIdentifier(layerName)

	for _, item := range layer.IntGrid {
		x := item.Position[0]
		y := item.Position[1]
		p := image.Point{
			X: x,
			Y: y,
		}

		data = append(data, p)
	}

	return data
}
