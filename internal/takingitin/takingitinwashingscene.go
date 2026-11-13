package takingitin

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/derekwaters/takingitin/internal/framework"
	"golang.org/x/image/colornames"
)

type TakingItInWashingScene struct {
	rightHand		*TakingItInHandObject
	leftHand		*TakingItInHandObject
}

func (scene *TakingItInWashingScene) Init(game *framework.PixelGameController) {
	scene.leftHand = NewTakingItInHandObject(game, "righthand", pixel.V(100, 100), pixel.KeyW, pixel.KeyS, pixel.KeyD, pixel.KeyA)
	scene.rightHand = NewTakingItInHandObject(game, "righthand", pixel.V(800, 100), pixel.KeyI, pixel.KeyK, pixel.KeyL, pixel.KeyJ)
}

func (scene *TakingItInWashingScene) Render(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	win.Clear(colornames.Green)
	scene.rightHand.Draw(win)
	scene.leftHand.Draw(win)
}

func (scene *TakingItInWashingScene) HandleKey(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	scene.rightHand.HandleKey(win)
	scene.leftHand.HandleKey(win)
}