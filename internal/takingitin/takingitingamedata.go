package takingitin

import (
	"github.com/gopxl/pixel/v2"
	"github.com/derekwaters/takingitin/internal/framework"
	"golang.org/x/image/colornames"
)

type TakingItInGameData struct {
	score			int
}

func NewTakingItInMenuData(controller *framework.PixelGameController) framework.PixelMenuData {
	options := []framework.PixelMenuSelection{
		framework.PixelMenuSelection_NewGame,
		framework.PixelMenuSelection_About,
		framework.PixelMenuSelection_Quit,
	}
	data := framework.PixelMenuData{
		BackgroundColor: colornames.Midnightblue,
		MenuTextColor: colornames.Whitesmoke,
		BackgroundImageTag: "menubg",
		MenuFontTag: "menu",
		TitleFontTag: "menutitle",
		MenuOptions: options,
		NewGameSceneName: "intro",
		AboutSceneName: "about",
	}
	return data
}

func (game TakingItInGameData) Init(controller *framework.PixelGameController) {
	controller.Log("--- Initialising TakingItIn ---")
	// Load Resources
	controller.Resources.AddFont("menu", "resources/fonts/notacumulus.ttf", 60)
	controller.Resources.AddFont("menutitle", "resources/fonts/notacumulus.ttf", 100)
	controller.Resources.AddPicture("menubg", "resources/pictures/stormbg.jpg")
	controller.Resources.AddPicture("righthand1", "resources/pictures/hand_r1.png")
	controller.Resources.AddPicture("righthand2", "resources/pictures/hand_r2.png")
	controller.Resources.AddPicture("righthand3", "resources/pictures/hand_r3.png")
	controller.Resources.AddPicture("righthand4", "resources/pictures/hand_r4.png")
	controller.Resources.AddPicture("lefthand1", "resources/pictures/hand_l1.png")
	controller.Resources.AddPicture("lefthand2", "resources/pictures/hand_l2.png")
	controller.Resources.AddPicture("lefthand3", "resources/pictures/hand_l3.png")
	controller.Resources.AddPicture("lefthand4", "resources/pictures/hand_l4.png")
	// Create Scenes
	menuData := NewTakingItInMenuData(controller)
	menuScene := framework.NewPixelMenuScene(menuData)
	var debugScene TakingItInDebugScene
	var washingScene TakingItInWashingScene
	controller.AddScene("menu", &menuScene)
	controller.AddScene("debug", &debugScene)
	controller.AddScene("about", &debugScene)
	controller.AddScene("intro", &washingScene)
	controller.AddScene("washing", &washingScene)
	controller.AddScene("end", &debugScene)
}

func (game TakingItInGameData) GetTitle() string {
	return "Taking It In!"
}

func (game TakingItInGameData) GetBounds() pixel.Rect {
	return pixel.R(0, 0, 1200, 960)
}

