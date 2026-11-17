package takingitin

import (
	"fmt"
	"github.com/gopxl/pixel/v2"
	"github.com/derekwaters/takingitin/internal/framework"
	"golang.org/x/image/colornames"
)

type TakingItInGameData struct {
	Score			int
	TimeRemaining	int

	PegsCollected	int
	PegsOnGround	int
	TotalPegs		int

	WashingCollected	int
	WashingOnGround		int
	TotalWashing		int
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
		MenuBgMusicName: "menu_loop",
		MenuNewOptionSoundName: "menu_button",
		MenuSelectOptionSoundName: "menu_button",
	}
	return data
}

func (game *TakingItInGameData) Init(controller *framework.PixelGameController) {
	controller.Log("--- Initialising TakingItIn ---")
	// Load Resources

	controller.Resources.AddFont("menu", "resources/fonts/notacumulus.ttf", 60)
	controller.Resources.AddFont("menutitle", "resources/fonts/notacumulus.ttf", 100)
	controller.Resources.AddFont("score", "resources/fonts/notacumulus.ttf", 45)
	controller.Resources.AddFont("keys", "resources/fonts/notacumulus.ttf", 25)
	controller.Resources.AddFont("debug", "resources/fonts/notacumulus.ttf", 15)

	var pictureData = map[string]string{
		"menubg": "resources/pictures/stormbg.jpg",
		"righthand1": "resources/pictures/hand_r1.png",
		"righthand2": "resources/pictures/hand_r2.png",
		"righthand3": "resources/pictures/hand_r3.png",
		"righthand4": "resources/pictures/hand_r4.png",
		"lefthand1": "resources/pictures/hand_l1.png",
		"lefthand2": "resources/pictures/hand_l2.png",
		"lefthand3": "resources/pictures/hand_l3.png",
		"lefthand4": "resources/pictures/hand_l4.png",

		"basket_back": "resources/pictures/basket_back.png",
		"bucket_back": "resources/pictures/bucket_back.png",
		"washing_bg": "resources/pictures/washing_bg.png",
	}
	var maskPictureData = map[string]string{		
		"peg_green": "resources/pictures/peg_green.png",
		"peg_red": "resources/pictures/peg_red.png",
		"peg_white": "resources/pictures/peg_white.png",
		"washing_shirt1": "resources/pictures/washing_shirt1.png",
		"washing_shirt2": "resources/pictures/washing_shirt2.png",
		"washing_undies1": "resources/pictures/washing_undies1.png",
		"basket_front": "resources/pictures/basket_front.png",
		"bucket_front": "resources/pictures/bucket_front.png",
	}
	controller.Resources.AddPictures(pictureData, false)
	controller.Resources.AddPictures(maskPictureData, true)

	var audioData = map[string]string{
		"menu_loop": "resources/sounds/rain_thunder_loop.wav",
		"menu_button": "resources/sounds/menu_button.wav",
		"washing1": "resources/sounds/washing_noise1.wav",
		"washing2": "resources/sounds/washing_noise2.wav",
		"washing3": "resources/sounds/washing_noise3.wav",
		"washing_bg": "resources/sounds/washing_bg.wav",
	}
	controller.Sounds.AddSounds(audioData)

	// Create Scenes
	menuData := NewTakingItInMenuData(controller)
	menuScene := framework.NewPixelMenuScene(menuData)
	var debugScene TakingItInDebugScene
	var introScene TakingItInIntroScene
	var washingScene TakingItInWashingScene
	introScene.SetGameData(game)
	washingScene.SetGameData(game)
	controller.AddScene("menu", &menuScene)
	controller.AddScene("debug", &debugScene)
	controller.AddScene("about", &debugScene)
	controller.AddScene("intro", &introScene)
	controller.AddScene("washing", &washingScene)
	controller.AddScene("end", &debugScene)
}

func (game *TakingItInGameData) GetTitle() string {
	return "Taking It In!"
}

func (game *TakingItInGameData) GetBounds() pixel.Rect {
	return pixel.R(0, 0, 1200, 960)
}

// Scoring
func (game *TakingItInGameData) AddCollectedPeg() {
	game.PegsCollected++
	game.Score += 5
	fmt.Println("score is ", game.Score)
}

func (game *TakingItInGameData) AddMissedPeg() {
	game.PegsOnGround++
	game.Score -= 2
	fmt.Println("score is ", game.Score)

}

func (game *TakingItInGameData) AddCollectedWashing() {
	game.WashingCollected++
	game.Score += 40

	fmt.Println("score is ", game.Score)
}

func (game *TakingItInGameData) AddMissedWashing() {
	game.WashingOnGround++
	game.Score -= 25
	fmt.Println("score is ", game.Score)
}
