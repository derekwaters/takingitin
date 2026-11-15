package takingitin

import (
	"fmt"
	"math"
	"math/rand"
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/imdraw"
	"github.com/gopxl/pixel/v2/ext/text"
	"github.com/derekwaters/takingitin/internal/framework"
	"golang.org/x/image/colornames"
)

const GROUND_LEVEL = 20

type TakingItInWashingScene struct {
	rightHand		*TakingItInHandObject
	leftHand		*TakingItInHandObject
	pegs			[]*TakingItInPegObject
	washing			[]*TakingItInWashingObject
	basket			*TakingItInBasketObject
	bucket			*TakingItInBucketObject
	scoreText		*text.Text
	timerText		*text.Text
	gameData		*TakingItInGameData
	imdraw			*imdraw.IMDraw

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
		pixel.V(200, 140), 
		pixel.KeyW, 
		pixel.KeyS, 
		pixel.KeyD, 
		pixel.KeyA,
		pixel.KeyLeftShift)
	scene.rightHand = NewTakingItInHandObject(
		game, 
		[]string{"righthand1", "righthand2", "righthand3", "righthand4"},
		pixel.V(1000, 140), 
		pixel.KeyI, 
		pixel.KeyK, 
		pixel.KeyL, 
		pixel.KeyJ,
		pixel.KeyRightShift)

	scene.SetupWashingLines(game)

	scene.basket = NewTakingItInBasketObject(
		game,
		"basket_back",
		"basket_front",
		pixel.V(200, 50))

	scene.bucket = NewTakingItInBucketObject(
		game,
		"bucket_back",
		"bucket_front",
		pixel.V(1000, 700))

	scene.scoreText = game.Resources.GenerateTextElement(
		pixel.V(10, 900), 
		"score", 
		colornames.White)
	scene.timerText = game.Resources.GenerateTextElement(
		pixel.V(800, 900), 
		"score", 
		colornames.White)

	// Debug
	scene.posText = game.Resources.GenerateTextElement(
		pixel.V(900, 200), 
		"debug", 
		colornames.Red)
}

func (scene *TakingItInWashingScene) SetupWashingLines(game *framework.PixelGameController) {
	itemsPerLine := 4

	for line := 0; line < 1/*5*/; line++ {
		yPos := float64(850 - (10 * line))
		xMin := float64(200 - (10 * line))
		xMax := float64(1000 + (10 * line))
		itemWidth := (xMax - xMin) / float64(itemsPerLine) 
		pegOffset := (itemWidth - 40) / 3

		for x := xMin; x < xMax; x += itemWidth {
			scene.AddWashingAndPegs(game, x + (itemWidth / 2), yPos, 0, "washing_shirt1", []float64{-pegOffset, pegOffset})
		}

		scene.imdraw = imdraw.New(nil)
		scene.imdraw.Color = colornames.Darkslategrey
		scene.imdraw.Push(pixel.V(xMin, yPos))
		scene.imdraw.Push(pixel.V(xMax, yPos))
		scene.imdraw.EndShape = imdraw.RoundEndShape
		scene.imdraw.Line(5.0)
	}
}

func (scene *TakingItInWashingScene) AddWashingAndPegs(
	game *framework.PixelGameController, 
	x float64, 
	y float64, 
	rotation float64,
	tag string, 
	pegOffsets []float64) {

	washing := NewWashingObject(game, tag, pixel.V(x, y), rotation)
	scene.washing = append(scene.washing, washing)
	scene.gameData.TotalWashing++

	for i := 0; i < len(pegOffsets); i++ {
		rotation := (rand.Float64() - 0.5) * math.Pi / 10
		peg := NewPegObject(game, PegColour_White, pixel.V(x + pegOffsets[i], y), rotation)
		washing.HeldByPeg(peg)
		peg.HoldingWashing(washing)
		scene.pegs = append(scene.pegs, peg)
		scene.gameData.TotalPegs++
	}
}

func (scene *TakingItInWashingScene) Render(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	// Draw the Background
	win.Clear(colornames.Green)

	// Draw the Back Of The Basket
	scene.basket.DrawBack(win)
	scene.bucket.DrawBack(win)

	// Draw the Line
	scene.imdraw.Draw(win)

	// Draw The Clothes
	for i := 0; i < len(scene.washing); i++ {
		scene.washing[i].Draw(win)
	}

	// Draw the Front Of The Basket
	scene.basket.DrawFront(win)

	// Draw the Pegs
	for i := 0; i < len(scene.pegs); i++ {
		scene.pegs[i].Draw(win)
	}
	
	// Draw the Front of the Peg Bucket
	scene.bucket.DrawFront(win)

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
	scene.rightHand.HandleKey(win, dt)
	scene.leftHand.HandleKey(win, dt)

	// Handle the hands gripping or otherwise
	scene.rightHand.HandleGripping(scene)
	scene.leftHand.HandleGripping(scene)

	for i := 0; i < len(scene.pegs); i++ {
		if scene.pegs[i].GetState() != PegState_InTheBucket {
			if scene.pegs[i].HandleKey(win, dt) {
				// Peg hit the ground - is it in a bucket?
				scene.PegHitGround(game, scene.pegs[i])
			}
			if scene.bucket.CheckInBucket(scene.pegs[i].GetPos()) {
				scene.pegs[i].Stop()
				scene.pegs[i].SetState(PegState_InTheBucket)
				scene.PegHitGround(game, scene.pegs[i])
			}
		}
	}

	for i := 0; i < len(scene.washing); i++ {
		if scene.washing[i].GetState() != WashingState_InTheBasket {
			if scene.washing[i].HandleKey(win, dt) {
				// Washing hit the ground - is it in a bucket?
				scene.WashingHitGround(game, scene.washing[i])
			}
			if scene.basket.CheckInBasket(scene.washing[i].GetPos()) {
				scene.washing[i].Stop()
				scene.washing[i].SetState(WashingState_InTheBasket)
				scene.WashingHitGround(game, scene.washing[i])
			}
		}
	}
}

func (scene *TakingItInWashingScene) PegHitGround(game *framework.PixelGameController, peg *TakingItInPegObject) {
	if (scene.bucket.CheckInBucket(peg.GetPos())) {
		// Nice!
		scene.gameData.AddCollectedPeg()
	} else {
		scene.gameData.AddMissedPeg()
	}
}

func (scene *TakingItInWashingScene) WashingHitGround(game *framework.PixelGameController, washing *TakingItInWashingObject) {
	sounds := []string{"washing1", "washing2", "washing3"}
	game.Sounds.Play(sounds[rand.Intn(len(sounds))])
	if (scene.basket.CheckInBasket(washing.GetPos())) {
		// Nice!
		scene.gameData.AddCollectedWashing()
	} else {
		scene.gameData.AddMissedWashing()
	}
}

func (scene *TakingItInWashingScene) FindGrabbedPeg(atPos pixel.Vec) (*TakingItInPegObject) {

	for i := 0; i < len(scene.pegs); i++ {
		if scene.pegs[i].CheckHit(atPos) {
			return scene.pegs[i]
		}
	}
	return nil
}

func (scene *TakingItInWashingScene) FindGrabbedWashing(atPos pixel.Vec) (*TakingItInWashingObject) {

	for i := 0; i < len(scene.washing); i++ {
		if scene.washing[i].CheckHit(atPos) {
			return scene.washing[i]
		}
	}
	return nil
}

func (scene *TakingItInWashingScene) GetLinePosY(x float64, line float64) (float64) {
	return (0.000125 * (float64(x) - 600.0) * (float64(x) - 600.0)) + 800.0
}

func (scene *TakingItInWashingScene) OnShowScene(game *framework.PixelGameController) {
	game.Sounds.StopAll()
	game.Sounds.Loop("washing_bg")
}