package eblib

import (
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
