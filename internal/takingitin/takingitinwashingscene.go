package takingitin

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/derekwaters/takingitin/internal/framework"
	"golang.org/x/image/colornames"
)

const MAX_SPEED = 4.0
const ACCEL = 0.5
const FRICTION = 0.3

type TakingItInWashingScene struct {
	rightHand		*pixel.Sprite
	rightHandPos	pixel.Matrix
	rightHandSpeed	pixel.Vec
}

func (scene *TakingItInWashingScene) Init(game *framework.PixelGameController) {
	pic := game.Resources.GetPicture("righthand")
	scene.rightHand = pixel.NewSprite(pic, pic.Bounds())
	scene.rightHandPos = pixel.IM.Scaled(pixel.V(0, 0), 0.4).Moved(pixel.V(100, 100));
	scene.rightHandSpeed = pixel.V(0, 0)
	game.Log("mat was ", scene.rightHandPos)
}

func (scene *TakingItInWashingScene) Render(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	win.Clear(colornames.Green)
	scene.rightHand.Draw(win, scene.rightHandPos)
}

func (scene *TakingItInWashingScene) HandleKey(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	if win.Pressed(pixel.KeyA) {
		scene.rightHandSpeed.X -= ACCEL
	} else if win.Pressed(pixel.KeyD) {
		scene.rightHandSpeed.X += ACCEL
	} else {
		if scene.rightHandSpeed.X < 0 {
			scene.rightHandSpeed.X += FRICTION
			if scene.rightHandSpeed.X > 0 {
				scene.rightHandSpeed.X = 0
			}
		} else if scene.rightHandSpeed.X > 0 {
			scene.rightHandSpeed.X -= FRICTION
			if scene.rightHandSpeed.X < 0 {
				scene.rightHandSpeed.X = 0
			}
		}
	}
	if win.Pressed(pixel.KeyS) {
		scene.rightHandSpeed.Y -= ACCEL
	} else if win.Pressed(pixel.KeyW) {
		scene.rightHandSpeed.Y += ACCEL
	} else {
		if scene.rightHandSpeed.Y < 0 {
			scene.rightHandSpeed.Y += FRICTION
			if scene.rightHandSpeed.Y > 0 {
				scene.rightHandSpeed.Y = 0
			}
		} else if scene.rightHandSpeed.Y > 0 {
			scene.rightHandSpeed.Y -= FRICTION
			if scene.rightHandSpeed.Y < 0 {
				scene.rightHandSpeed.Y = 0
			}
		}
	}
	// Clamp the speed
	if scene.rightHandSpeed.X < (-1.0 * MAX_SPEED) {
		scene.rightHandSpeed.X = -1.0 * MAX_SPEED
	} else if scene.rightHandSpeed.X > MAX_SPEED {
		scene.rightHandSpeed.X = MAX_SPEED
	}
	if scene.rightHandSpeed.Y < (-1.0 * MAX_SPEED) {
		scene.rightHandSpeed.Y = -1.0 * MAX_SPEED
	} else if scene.rightHandSpeed.Y > MAX_SPEED {
		scene.rightHandSpeed.Y = MAX_SPEED
	}
	// Now move the hand
	scene.rightHandPos = scene.rightHandPos.Moved(scene.rightHandSpeed)
}