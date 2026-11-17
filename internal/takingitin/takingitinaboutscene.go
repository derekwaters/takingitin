package takingitin

import (
	"fmt"
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/text"
	"github.com/derekwaters/takingitin/internal/framework"
	"golang.org/x/image/colornames"
)

type TakingItInAboutScene struct {
	about			*text.Text
	bgSprite		*pixel.Sprite
	bgBounds		pixel.Rect
}

func (scene *TakingItInAboutScene) Init(game *framework.PixelGameController) {
	scene.about = game.Resources.GenerateTextElement(
		pixel.V(50, 800), 
		"menu", 
		colornames.White)

	scene.about.Clear()
	fmt.Fprintf(scene.about, 
`Welcome to "Taking It In!", 
a game written for the 
itch.io "Hand Crafted Jam".
Thanks to Dylan for the jam!

Idea By: Angus
Tactile By: Angus' actual hands
Hand Crafted By: Derek
Background Noises: Nox Sound
Sound Effects: Film Cow
Fonts: Chequered Ink`)

	bgImage := game.Resources.GetPicture("menubg")
	if bgImage != nil {
		scene.bgBounds = bgImage.Bounds()
		scene.bgSprite = pixel.NewSprite(bgImage, scene.bgBounds)
	}
}

func (scene *TakingItInAboutScene) Render(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	// Draw the Background
	if scene.bgSprite != nil {
		scaleFactor := win.Bounds().H() / scene.bgBounds.H()
		scene.bgSprite.Draw(win, pixel.IM.Scaled(pixel.V(0.0, 0.0), scaleFactor).Moved(win.Bounds().Center()))
	}

	// Draw the Text
	scene.about.Draw(win, pixel.IM)
}

func (scene *TakingItInAboutScene) HandleKey(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	if win.JustPressed(pixel.KeyEnter) ||
		win.JustPressed(pixel.KeyEscape) {
		game.SetScene("menu")
	}
}

func (scene *TakingItInAboutScene) OnShowScene(game *framework.PixelGameController) {
}
