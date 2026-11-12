package takingitin

import (
	"github.com/derekwaters/takingitin/internal/framework"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"golang.org/x/image/colornames"
)


type TakingItInDebugScene struct {
}

func (scene TakingItInDebugScene) Init(game *framework.PixelGameController) {
	
}

func (scene TakingItInDebugScene) Render(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	win.Clear(colornames.Blue)
}

func (scene TakingItInDebugScene) HandleKey(game *framework.PixelGameController, win *opengl.Window, dt float64) {
}