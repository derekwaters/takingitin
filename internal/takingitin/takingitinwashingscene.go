package takingitin

import (
	"github.com/derekwaters/takingitin/internal/framework"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"golang.org/x/image/colornames"
)


type TakingItInWashingScene struct {
	
}

func (scene TakingItInWashingScene) Init(game *framework.PixelGameController) {
	
}

func (scene TakingItInWashingScene) Render(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	win.Clear(colornames.Green)
}

func (scene TakingItInWashingScene) HandleKey(game *framework.PixelGameController, win *opengl.Window, dt float64) {
}