package takingitin

import (
	"fmt"
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/imdraw"
	"github.com/gopxl/pixel/v2/ext/text"
	"github.com/derekwaters/takingitin/internal/framework"
	"golang.org/x/image/colornames"
)

type TakingItInIntroScene struct {
	rightHand		*TakingItInHandObject
	leftHand		*TakingItInHandObject
	instructions	*text.Text
	keys			[]*text.Text
	imdraw			*imdraw.IMDraw

	gameData		*TakingItInGameData
}


func (scene *TakingItInIntroScene) SetGameData(gameData *TakingItInGameData) {
	scene.gameData = gameData
}

func (scene *TakingItInIntroScene) Init(game *framework.PixelGameController) {
	scene.leftHand = NewTakingItInHandObject(
		game, 
		[]string{"lefthand1"}, 
		pixel.V(200, 140), 
		pixel.KeyW, 
		pixel.KeyS, 
		pixel.KeyD, 
		pixel.KeyA,
		pixel.KeyLeftShift)
	scene.rightHand = NewTakingItInHandObject(
		game, 
		[]string{"righthand1"},
		pixel.V(1000, 140), 
		pixel.KeyI, 
		pixel.KeyK, 
		pixel.KeyL, 
		pixel.KeyJ,
		pixel.KeyRightShift)
		
	scene.instructions = game.Resources.GenerateTextElement(
		pixel.V(100, 800), 
		"score", 
		colornames.White)

	scene.imdraw = imdraw.New(nil)
	scene.imdraw.Color = colornames.White
	// Left Controls
	rects := []float64{
		180,  200, 220,  240, 
		180,   40, 220,   80,
		100,  120, 140,  160,
		260,  120, 300,  160,
		980,  200, 1020, 240, 
		980,   40, 1020,  80,
		900,  120, 940,  160,
		1060, 120, 1100, 160,
		240,   40, 340,   80,
		1040,  40, 1140,  80,
	}
	labels := []string{"W", "S", "A", "D", "I", "K", "J", "L", "LShift", "RShift"}
	for i := 0; i < len(rects); i += 4 {
		// textPos := pixel.V((rects[i] + rects[i + 2]) / 2, (rects[i + 1] + rects[i + 3]) / 2)
		textPos := pixel.V(rects[i] + 10, rects[i + 1] + 10)
		newText := game.Resources.GenerateTextElement(
			textPos, 
			"keys", 
			colornames.White)
		newText.Clear()
		fmt.Fprintf(newText, "%s", labels[i / 4])
		
		scene.keys = append(scene.keys, newText)

		scene.imdraw.Push(pixel.V(rects[i], rects[i + 1]))
		scene.imdraw.Push(pixel.V(rects[i + 2], rects[i + 3]))
		scene.imdraw.Rectangle(4.0)
	}
}

func (scene *TakingItInIntroScene) Render(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	// Draw the Background
	win.Clear(colornames.Darkslategrey)

	// Draw the Hands
	scene.rightHand.Draw(win)
	scene.leftHand.Draw(win)

	scene.imdraw.Draw(win)

	// Draw the Text
	scene.instructions.Clear()
	fmt.Fprintf(scene.instructions, "Uh, oh, it looks like it's about to rain...\nAnd I've got washing on the line...\nTime to...\n\nTake It In!")
	scene.instructions.Draw(win, pixel.IM)

	for i := 0; i < len(scene.keys); i++ {
		scene.keys[i].Draw(win, pixel.IM)
	}
}

func (scene *TakingItInIntroScene) HandleKey(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	if win.JustPressed(pixel.KeyEnter) {
		game.SetScene("washing")
	}
}

func (scene *TakingItInIntroScene) OnShowScene(game *framework.PixelGameController) {
}
