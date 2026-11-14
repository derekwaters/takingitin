package takingitin

import (
	"math/rand"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/derekwaters/takingitin/internal/framework"
)

const PEG_SCALE = 0.2

type PegState int
type PegColour int

const (
	PegState_OnTheLine = iota
	PegState_Falling
	PegState_OnTheGround
	PegState_InTheBasket
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
}

func NewPegObject(
	game *framework.PixelGameController,
	colour PegColour,
	startPos pixel.Vec) (*TakingItInPegObject) {

	physObj := NewTakingItInPhysicsObject(
		game, 
		[]string{PegColourName[colour]},
		PEG_SCALE,
		startPos,
		pixel.ZV,
		0,
		0,
		PhysicsGroundCollisionMode_BounceDirect)

	peg := TakingItInPegObject{
		object: physObj,
		pegState: PegState_OnTheLine,
		pegColour: colour,
	}
	return &peg
}

func (peg *TakingItInPegObject) Draw(win *opengl.Window) {
	peg.object.Draw(win)
}

func (peg *TakingItInPegObject) HandleKey(win *opengl.Window, dt float64) {
	peg.object.UpdatePos(dt, 20)
}

func (peg *TakingItInPegObject) Unpeg() {
	peg.pegState = PegState_Falling
	peg.object.Fall(0.1 * (rand.Float64() - 0.5))
}