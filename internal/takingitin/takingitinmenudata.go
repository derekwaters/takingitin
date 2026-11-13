package takingitin

import (
	"github.com/derekwaters/takingitin/internal/framework"
	"golang.org/x/image/colornames"
)

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
