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
	pegs			[]*TakingItInPegObject
}

func (scene *TakingItInWashingScene) Init(game *framework.PixelGameController) {
	scene.leftHand = NewTakingItInHandObject(
		game, 
		[]string{"lefthand1", "lefthand2", "lefthand3", "lefthand4"}, 
		pixel.V(100, 100), 
		pixel.KeyW, 
		pixel.KeyS, 
		pixel.KeyD, 
		pixel.KeyA,
		pixel.KeyLeftShift)
	scene.rightHand = NewTakingItInHandObject(
		game, 
		[]string{"righthand1", "righthand2", "righthand3", "righthand4"},
		pixel.V(800, 100), 
		pixel.KeyI, 
		pixel.KeyK, 
		pixel.KeyL, 
		pixel.KeyJ,
		pixel.KeyRightShift)
	scene.pegs = append(scene.pegs, NewPegObject(game, PegColour_Green, pixel.V(200, 700)))
	scene.pegs = append(scene.pegs, NewPegObject(game, PegColour_Red, pixel.V(300, 700)))
	scene.pegs = append(scene.pegs, NewPegObject(game, PegColour_White, pixel.V(400, 700)))
}

func (scene *TakingItInWashingScene) Render(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	win.Clear(colornames.Green)
	scene.rightHand.Draw(win)
	scene.leftHand.Draw(win)
	for i := 0; i < len(scene.pegs); i++ {
		scene.pegs[i].Draw(win)
	}
}

func (scene *TakingItInWashingScene) HandleKey(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	scene.rightHand.HandleKey(win)
	scene.leftHand.HandleKey(win)
	for i := 0; i < len(scene.pegs); i++ {
		if win.JustPressed(pixel.KeyB) {
			scene.pegs[i].Unpeg()
		}
		scene.pegs[i].HandleKey(win, dt)
	}
}