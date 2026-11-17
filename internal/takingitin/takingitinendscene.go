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

type TakingItInEndScene struct {
	scores			*text.Text
	imdraw			*imdraw.IMDraw

	gameData		*TakingItInGameData
}


func (scene *TakingItInEndScene) SetGameData(gameData *TakingItInGameData) {
	scene.gameData = gameData
}

func (scene *TakingItInEndScene) Init(game *framework.PixelGameController) {
	scene.scores = game.Resources.GenerateTextElement(
		pixel.V(100, 800), 
		"score", 
		colornames.White)
}

func (scene *TakingItInEndScene) Render(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	// Draw the Background
	win.Clear(colornames.Darkslategrey)

	// Draw the Text
	scene.scores.Clear()
	fmt.Fprintf(scene.scores, "Your score: %d\n\nYou collected %d pieces of washing out of %d\nYou collected %d pegs out of %d", 
		scene.gameData.Score,
		scene.gameData.WashingCollected,
		scene.gameData.TotalWashing,
		scene.gameData.PegsCollected,
		scene.gameData.TotalPegs)
	scene.scores.Draw(win, pixel.IM)
}

func (scene *TakingItInEndScene) HandleKey(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	if win.JustPressed(pixel.KeyEnter) ||
		win.JustPressed(pixel.KeyEscape) {
		game.SetScene("menu")
	}
}

func (scene *TakingItInEndScene) OnShowScene(game *framework.PixelGameController) {
	game.Sounds.StopAll()
}
