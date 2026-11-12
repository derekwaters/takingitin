package takingitin

import (
	"github.com/derekwaters/takingitin/internal/framework"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"golang.org/x/image/colornames"
)


type TakingItInRoomScene struct {
}

func (scene TakingItInRoomScene) Init(game *framework.PixelGameController) {
	
}

func (scene TakingItInRoomScene) Render(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	win.Clear(colornames.Green)
}

func (scene TakingItInRoomScene) HandleKey(game *framework.PixelGameController, win *opengl.Window, dt float64) {
}