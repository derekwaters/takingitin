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

const GROUND_LEVEL = 10

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
	bgSprite		*pixel.Sprite
	bgBounds		pixel.Rect

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
		pixel.V(200, 10))

	scene.bucket = NewTakingItInBucketObject(
		game,
		"bucket_back",
		"bucket_front",
		pixel.V(1120, 610))

	scene.scoreText = game.Resources.GenerateTextElement(
		pixel.V(40, 900), 
		"score", 
		colornames.Darkslategrey)
	scene.timerText = game.Resources.GenerateTextElement(
		pixel.V(800, 900), 
		"score", 
		colornames.White)

	// Debug
	scene.posText = game.Resources.GenerateTextElement(
		pixel.V(900, 200), 
		"debug", 
		colornames.Red)

	bgImage := game.Resources.GetPicture("washing_bg")
	if bgImage != nil {
		scene.bgBounds = bgImage.Bounds()
		scene.bgSprite = pixel.NewSprite(bgImage, scene.bgBounds)
	}
}

type washingDef struct {
	label	string
	width	float64
	pegs	[]float64
}

func newWashingDef(label string, width float64, pegs []float64) (washingDef) {
	return washingDef{
		label: label,
		width: width,
		pegs: pegs,
	}
}

func (scene *TakingItInWashingScene) SetupWashingLines(game *framework.PixelGameController) {

	clothesTypes := make([]washingDef, 0, 3)
	clothesTypes = append(clothesTypes, newWashingDef("washing_shirt1", 400.0, []float64{-160.0, 150.0}))
	clothesTypes = append(clothesTypes, newWashingDef("washing_shirt2", 400.0, []float64{-142.0, 157.0}))
	clothesTypes = append(clothesTypes, newWashingDef("washing_undies1", 160.0, []float64{21.0}))

	game.Log(clothesTypes)

	spaceForBucket := 80.0
	spaceBetweenItems := 20.0

	for line := 0; line < 1/*5*/; line++ {
		yPos := float64(720 - (10 * line))
		xMin := float64(0 + (10 * line))
		xMax := float64(1200 - (10 * line))
		washingScale := 0.4 - (0.02 * float64(line))

		for x := xMin + spaceBetweenItems; x < xMax - spaceForBucket; {
			itemType := rand.Intn(len(clothesTypes))
			item := clothesTypes[itemType]

			theItemWidth := item.width * washingScale;

			if x + theItemWidth < xMax - spaceForBucket {

				var pegOffsets []float64
				for q := 0; q < len(item.pegs); q++ {
					pegOffsets = append(pegOffsets, item.pegs[q] * washingScale)
				}
				game.Log("Adding a ", item.label, "with width", theItemWidth, itemType)
				scene.AddWashingAndPegs(game, x + theItemWidth / 2, yPos, 0, item.label, pegOffsets)
			}

			x += theItemWidth + spaceBetweenItems
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

	washing := NewWashingObject(game, tag, pixel.V(x, y), rotation, 0.4)
	scene.washing = append(scene.washing, washing)
	scene.gameData.TotalWashing++

	for i := 0; i < len(pegOffsets); i++ {
		rotation := (rand.Float64() - 0.5) * math.Pi / 10
		pegColour := PegColour(rand.Intn(int(PegColour_Green) + 1))
		peg := NewPegObject(game, pegColour, pixel.V(x + pegOffsets[i], y), rotation)
		washing.HeldByPeg(peg)
		peg.HoldingWashing(washing)
		scene.pegs = append(scene.pegs, peg)
		scene.gameData.TotalPegs++
	}
}

func (scene *TakingItInWashingScene) Render(game *framework.PixelGameController, win *opengl.Window, dt float64) {
	// Draw the Background
	// win.Clear(colornames.Green)

	if scene.bgSprite != nil {
		scaleFactor := win.Bounds().H() / scene.bgBounds.H()
		scene.bgSprite.Draw(win, pixel.IM.Scaled(pixel.V(0.0, 0.0), scaleFactor).Moved(win.Bounds().Center()))
	}

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
			if scene.pegs[i].HandleKey(win, dt) || 
				scene.bucket.CheckInBucket(scene.pegs[i].GetPos()) {
				// Peg hit the ground - is it in a bucket?
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
		}
	}
}

func (scene *TakingItInWashingScene) PegHitGround(game *framework.PixelGameController, peg *TakingItInPegObject) {
	if (scene.bucket.CheckInBucket(peg.GetPos())) {
		// Nice!
		peg.Stop()
		peg.SetState(PegState_InTheBucket)
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
		washing.Stop()
		washing.SetState(WashingState_InTheBasket)
		scene.gameData.AddCollectedWashing()
	} else {
		washing.Stop()
		washing.SetState(WashingState_OnTheGround)
		scene.gameData.AddMissedWashing()
	}

	// Check to see if the game is over
	scene.CheckAllWashing(game)
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

func (scene *TakingItInWashingScene) CheckAllWashing(game *framework.PixelGameController) {
	for i := 0; i < len(scene.washing); i++ {
		game.Log("Checking all the washing, item ", i, "state", scene.washing[i].GetState())

		if scene.washing[i].GetState() != WashingState_OnTheGround &&
			scene.washing[i].GetState() != WashingState_InTheBasket {
			return
		}
	}

	// We're all done
	game.Log("All Washing Sorted")
	game.SetScene("end")
}