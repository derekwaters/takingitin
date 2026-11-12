package framework

import (
	// "github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
)

type PixelGameScene interface {
	Init(game *PixelGameController)
	Render(game *PixelGameController, win *opengl.Window, dt float64)
	HandleKey(game *PixelGameController, win *opengl.Window, dt float64)
}
