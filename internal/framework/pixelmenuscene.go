package framework

import (
	"fmt"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/text"
)


type PixelMenuScene struct {
	data					PixelMenuData
	bgSprite				*pixel.Sprite
	bgBounds				pixel.Rect
	currentMenuSelection	int
	titleText				*text.Text
	menuText				*text.Text
}

func NewPixelMenuScene(data PixelMenuData) (PixelMenuScene) {
	return PixelMenuScene{
		data: data,
		bgSprite: nil,
	}
}

func (scene *PixelMenuScene) Init(game *PixelGameController) {
	if len(scene.data.BackgroundImageTag) > 0 {
		bgImage := game.Resources.GetPicture(scene.data.BackgroundImageTag)
		if bgImage != nil {
			scene.bgBounds = bgImage.Bounds()
			scene.bgSprite = pixel.NewSprite(bgImage, scene.bgBounds)
		}
	}
	scene.currentMenuSelection = 0
	scene.titleText = game.Resources.GenerateTextElement(
		pixel.V(60, 700), 
		scene.data.TitleFontTag, 
		scene.data.MenuTextColor)
	scene.menuText = game.Resources.GenerateTextElement(
		pixel.V(60, 300), 
		scene.data.MenuFontTag,
		scene.data.MenuTextColor)
}

func (scene *PixelMenuScene) Render(game *PixelGameController, win *opengl.Window, dt float64) {
	win.Clear(scene.data.BackgroundColor)

	if scene.bgSprite != nil {
		scaleFactor := win.Bounds().H() / scene.bgBounds.H()
		scene.bgSprite.Draw(win, pixel.IM.Scaled(pixel.V(0.0, 0.0), scaleFactor).Moved(win.Bounds().Center()))
	}

	// Now draw the menu text

	totalText := ""
	for i := 0; i < len(scene.data.MenuOptions); i++ {
		if i == scene.currentMenuSelection {
			totalText += "> "
		}
		totalText += PixelMenuSelectionName[scene.data.MenuOptions[i]] + "\n"
	}
	scene.titleText.Clear()
	scene.menuText.Clear()
	fmt.Fprintf(scene.titleText, "%s", game.gameData.GetTitle())
	fmt.Fprintf(scene.menuText, "%s", totalText)
	scene.titleText.Draw(win, pixel.IM)
	scene.menuText.Draw(win, pixel.IM)
}

func (scene *PixelMenuScene) HandleKey(game *PixelGameController, win *opengl.Window, dt float64) {
	if win.JustPressed(pixel.KeyUp) {
		scene.currentMenuSelection--
		if scene.currentMenuSelection < 0 {
			scene.currentMenuSelection = 0
		}
	}
	if win.JustPressed(pixel.KeyDown) {
		scene.currentMenuSelection++
		if scene.currentMenuSelection >= len(scene.data.MenuOptions) {
			scene.currentMenuSelection = len(scene.data.MenuOptions) - 1
		}
	}
	if win.JustPressed(pixel.KeyEnter) {
		switch scene.data.MenuOptions[scene.currentMenuSelection] {
		case PixelMenuSelection_NewGame:
			game.SetScene(scene.data.NewGameSceneName)
		case PixelMenuSelection_About:
			game.SetScene(scene.data.AboutSceneName)
		case PixelMenuSelection_Options:
			// TODO:
		case PixelMenuSelection_Quit:
			win.SetClosed(true)
			return;
		}
	}
	if win.JustPressed(pixel.KeyEscape) {
		win.SetClosed(true)
	}
}