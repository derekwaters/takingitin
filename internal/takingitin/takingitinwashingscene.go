package takingitin

import (
	"fmt"
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/text"
	"github.com/derekwaters/takingitin/internal/framework"
	"golang.org/x/image/colornames"
)

type TakingItInWashingScene struct {
	rightHand		*TakingItInHandObject
	leftHand		*TakingItInHandObject
	pegs			[]*TakingItInPegObject
	basket			*TakingItInBasketObject
	scoreText		*text.Text
	timerText		*text.Text
	gameData		*TakingItInGameData

	// DEBUG
	posText			*text.Text
}


func (scene *TakingItInWashingScene) SetGameData(gameData *TakingItInGameData) {
	scene.gameData = gameData
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
	scene.basket = NewTakingItInBasketObject(
		game,
		"basket_back",
		"basket_front",
		pixel.V(200, 50))


	scene.scoreText = game.Resources.GenerateTextElement(
		pixel.V(10, 800), 
		"score", 
		colornames.White)
	scene.timerText = game.Resources.GenerateTextElement(
		pixel.V(800, 800), 
		"score", 
		colornames.White)

	// Debug
	scene.posText = game.Resources.GenerateTextElement(
		pixel.V(900, 900), 
		"debug", 
		colornames.Red)
}

func (scene *TakingItInWashingScene) Render(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	// Draw the Background
	win.Clear(colornames.Green)

	// Draw the Back Of The Basket
	scene.basket.DrawBack(win)

	// Draw The Clothes

	// Draw the Front Of The Basket
	scene.basket.DrawFront(win)

	// Draw the Pegs
	for i := 0; i < len(scene.pegs); i++ {
		scene.pegs[i].Draw(win)
	}

	// Draw the Hands
	scene.rightHand.Draw(win)
	scene.leftHand.Draw(win)

	// Draw the Scores
	scene.scoreText.Clear()
	fmt.Fprintf(scene.scoreText, "Score: %d", scene.gameData.Score)
	scene.scoreText.Draw(win, pixel.IM)
	scene.timerText.Clear()
	fmt.Fprintf(scene.timerText, "Time: %d", scene.gameData.TimeRemaining)
	scene.timerText.Draw(win, pixel.IM)

	// Debug
	scene.posText.Clear()
	fmt.Fprintf(scene.posText, "Mouse: [%d, %d]", int(win.MousePosition().X), int(win.MousePosition().Y))
	scene.posText.Draw(win, pixel.IM)
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