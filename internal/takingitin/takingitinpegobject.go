package takingitin

import (
	"math"
	"math/rand"

	"github.com/gopxl/pixel/v2"
//	"github.com/gopxl/pixel/v2/ext/imdraw"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/derekwaters/takingitin/internal/framework"
)

const PEG_SCALE = 0.2
const PEG_SEND_FLYING_MAX_SPEED = 500.0

type PegState int
type PegColour int

const (
	PegState_OnTheLine = iota
	PegState_Falling
	PegState_OnTheGround
	PegState_InTheBucket
)

const (
	PegColour_Red = iota
	PegColour_White
	PegColour_Green
)


var PegColourName = map[PegColour]string{
	PegColour_Red: "peg_red",
	PegColour_White: "peg_white",
	PegColour_Green: "peg_green",
}


type TakingItInPegObject struct {
	object			*TakingItInPhysicsObject
	pegState		PegState
	pegColour		PegColour
	mask			*framework.PixelPictureMask
	holding			*TakingItInWashingObject
}

func NewPegObject(
	game *framework.PixelGameController,
	colour PegColour,
	startPos pixel.Vec,
	startRotation float64) (*TakingItInPegObject) {

	physObj := NewTakingItInPhysicsObject(
		game, 
		[]string{PegColourName[colour]},
		PEG_SCALE,
		startPos,
		pixel.ZV,
		startRotation,
		0,
		PhysicsGroundCollisionMode_BounceRandom,
		[]string{"peg1", "peg2", "peg3", "peg4"})
	peg := TakingItInPegObject{
		object: physObj,
		pegState: PegState_OnTheLine,
		pegColour: colour,
	}

	peg.mask = game.Resources.GetMask(PegColourName[colour])
	return &peg
}

func (peg *TakingItInPegObject) Draw(win *opengl.Window) {

	peg.object.Draw(win)
	
	// DEBUG
	// Check whether the mouse is on this item
	// if peg.CheckHit(win.MousePosition()) {
	// 	imd := imdraw.New(nil)
	// 	imd.Color = pixel.RGB(0, 1, 0)
	// 	imd.Push(peg.GetPos())
	// 	imd.Circle(3, 0)
	// 	imd.Draw(win)
	// }
}

func (peg *TakingItInPegObject) HandleKey(win *opengl.Window, dt float64) bool {
	return peg.object.UpdatePos(dt, GROUND_LEVEL)
}

func (peg *TakingItInPegObject) HoldingWashing(washing *TakingItInWashingObject) {
	peg.holding = washing
}

func (peg *TakingItInPegObject) Unpeg() {
	// Check the attached washing
	if peg.holding != nil {
		peg.holding.RemovePeg(peg)
	}
}

func (peg *TakingItInPegObject) SendFlying() {
	peg.pegState = PegState_Falling
	peg.object.SetSpeed(pixel.ZV.Unit().Rotated(rand.Float64() * math.Pi * 2).Scaled(rand.Float64() * PEG_SEND_FLYING_MAX_SPEED))
	peg.object.Fall(0.1 * (rand.Float64() - 0.5))
}

func (peg *TakingItInPegObject) Drop(horizSpeed float64) {
	peg.pegState = PegState_Falling
	peg.object.SetSpeed(pixel.V(horizSpeed, 0))
	peg.object.Fall(0.1 * (rand.Float64() - 0.5))
}

func (peg *TakingItInPegObject) GetPos() (pixel.Vec) {
	return peg.object.GetPos()
}

func (peg *TakingItInPegObject) SetPos(newPos pixel.Vec) {
	peg.object.SetPos(newPos)
}

func (peg *TakingItInPegObject) Stop() {
	peg.object.HitGround()
}

func (peg *TakingItInPegObject) GetState() (PegState) {
	return peg.pegState
}

func (peg *TakingItInPegObject) SetState(newState PegState) {
	peg.pegState = newState
}

// Note, this pos is in game screen coordinates
func (peg *TakingItInPegObject) CheckHit(atPos pixel.Vec) bool {

	checkInMask := atPos.Sub(peg.GetPos()).Scaled(1.0 / PEG_SCALE).Rotated(-peg.object.GetRotation()).Add(peg.mask.GetRect().Center())
	return peg.mask.GetMask(checkInMask)
}