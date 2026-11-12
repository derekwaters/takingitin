package framework

import (
	"fmt"
	"time"

	// "github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	// "github.com/gopxl/pixel/v2/ext/text"
)

type PixelGameController struct {
	
	// Game Data
	gameData					PixelGameData

	// Game Management Data
	Scenes						map[string]PixelGameScene
	CurrentScene				string

	// Pixel Data
	win							*opengl.Window
	frames						int
	second						<-chan time.Time
	DebugMode					bool

	// textAtlas					*text.Atlas
	// debugAtlas					*text.Atlas
	Resources					*PixelResourceController
}

func (game *PixelGameController) Launch(gameData PixelGameData) {
	game.Log(":: Starting Pixel")
	game.gameData = gameData
	runGame = game
	opengl.Run(run)
	game.Log(":: Finishing Pixel")
}

func (game *PixelGameController) AddScene(name string, scene PixelGameScene) {
	game.Scenes[name] = scene
	if len(game.Scenes) == 1 {
		game.SetScene(name)
	}
}

func (game *PixelGameController) SetScene(name string) {
	_, exist := game.Scenes[name]
	if exist {
		game.Log(":: Scene Changed to ", name)
		game.CurrentScene = name
	}
}


func (game *PixelGameController) Log(msg ...interface{}) {
	if (game.DebugMode) {
		fmt.Println(msg)
	}
}


func (game *PixelGameController) Init() {
	game.Resources = new(PixelResourceController)
	game.Resources.Init(game)
	game.gameData.Init(game)
	// Init all the scenes
	for _, scene := range game.Scenes {
		scene.Init(game)
	}
}

func (game *PixelGameController) DoRun() {
	game.Scenes = make(map[string]PixelGameScene)

	// Initiate the window
	cfg := opengl.WindowConfig{
		Title: game.gameData.GetTitle(),
		Bounds: game.gameData.GetBounds(),
		VSync: true,
	}

	win, err := opengl.NewWindow(cfg)
	game.win = win
	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	win.SetSmooth(true)

	game.Log(":: About to Init Game")
	game.Init()
	game.Log(":: Init Game Finished")

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()
		
		game.Scenes[game.CurrentScene].HandleKey(game, win, dt)
		game.Scenes[game.CurrentScene].Render(game, win, dt)
		
		win.Update()
	}
}



var runGame	*PixelGameController

func run() {
	// Initiate the Scene list
	runGame.DoRun()
}