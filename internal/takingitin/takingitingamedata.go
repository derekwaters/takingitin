package takingitin

import (
	"github.com/gopxl/pixel/v2"
	"github.com/derekwaters/takingitin/internal/framework"
)

type TakingItInGameData struct {
	score			int
}

func (game TakingItInGameData) Init(controller *framework.PixelGameController) {
	controller.Log("--- Initialising TakingItIn ---")
	// Load Resources
	controller.Resources.AddFont("menu", "resources/fonts/ancient-modern-tales.ttf", 60)
	controller.Resources.AddFont("menutitle", "resources/fonts/ancient-modern-tales.ttf", 100)
	controller.Resources.AddPicture("menubg", "resources/pictures/menubg.jpg")
	// Create Scenes
	menuData := NewTakingItInMenuData(controller)
	menuScene := framework.NewPixelMenuScene(menuData)
	var debugScene TakingItInDebugScene
	var roomScene TakingItInRoomScene
	controller.AddScene("menu", &menuScene)
	controller.AddScene("debug", &debugScene)
	controller.AddScene("room", &roomScene)

}

func (game TakingItInGameData) GetTitle() string {
	return "Taking It In!"
}

func (game TakingItInGameData) GetBounds() pixel.Rect {
	return pixel.R(0, 0, 1200, 960)
}

