package takingitin

import (
	//"math/rand"

	"github.com/gopxl/pixel/v2"
//	"github.com/gopxl/pixel/v2/ext/imdraw"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/derekwaters/takingitin/internal/framework"
)

type WashingState int

const (
	WashingState_OnTheLine = iota
	WashingState_Falling
	WashingState_OnTheGround
	WashingState_InTheBasket
)

type TakingItInWashingObject struct {
	object			*TakingItInPhysicsObject
	state			WashingState
	mask			*framework.PixelPictureMask
	heldByPegs		[]*TakingItInPegObject
	scale			float64
}

func NewWashingObject(
	game *framework.PixelGameController,
	item string,
	startPos pixel.Vec,
	startRotation float64,
	scale float64) (*TakingItInWashingObject) {

	pic := game.Resources.GetPicture(item)
	yOff := (pic.Bounds().H() / 2) * scale
	startPos.Y -= yOff

	physObj := NewTakingItInPhysicsObject(
		game, 
		[]string{item},
		scale,
		startPos,
		pixel.ZV,
		startRotation,
		0,
		PhysicsGroundCollisionMode_Stop)

	washing := TakingItInWashingObject{
		object: physObj,
		state: WashingState_OnTheLine,
		scale: scale,
	}
	washing.mask = game.Resources.GetMask(item)
	return &washing
}

func (washing *TakingItInWashingObject) Draw(win *opengl.Window) {

	washing.object.Draw(win)
	
	// DEBUG
	// Check whether the mouse is on this item
	// imd := imdraw.New(nil)
	// if washing.CheckHit(win.MousePosition()) {		
	// 	imd.Color = pixel.RGB(0, 1, 1)
	// } else {
	// 	imd.Color = pixel.RGB(1, 1, 0)
	// }
	// imd.Push(washing.GetPos())
	// imd.Circle(5, 0)
	// imd.Draw(win)
}

func (washing *TakingItInWashingObject) HandleKey(win *opengl.Window, dt float64) bool {
	return washing.object.UpdatePos(dt, GROUND_LEVEL)
}

func (washing *TakingItInWashingObject) HeldByPeg(peg *TakingItInPegObject) {
	washing.heldByPegs = append(washing.heldByPegs, peg)
}

// This send all the pegs flying
func (washing *TakingItInWashingObject) Unpeg() {
	for i := 0; i < len(washing.heldByPegs); i++ {
		washing.heldByPegs[i].SendFlying()
	}
	washing.state = WashingState_Falling
}

// This removes a single peg which may cause the washing to fall
func (washing *TakingItInWashingObject) RemovePeg(peg *TakingItInPegObject) {
	removeIndex := -1
	for i := 0; i < len(washing.heldByPegs) && (removeIndex == -1); i++ {
		if washing.heldByPegs[i] == peg {
			removeIndex = i
		}
	}
	if removeIndex >= 0 {
		peg.HoldingWashing(nil)
		washing.heldByPegs = append(washing.heldByPegs[:removeIndex], washing.heldByPegs[removeIndex + 1:]...)
	}

	if len(washing.heldByPegs) == 0 {
		washing.Drop(0.0)
	}
}


func (washing *TakingItInWashingObject) Drop(horizSpeed float64) {
	washing.state = WashingState_Falling
	washing.object.SetSpeed(pixel.V(horizSpeed, 0))
	washing.object.Fall(0.0)
}

func (washing *TakingItInWashingObject) GetPos() (pixel.Vec) {
	return washing.object.GetPos()
}

func (washing *TakingItInWashingObject) SetPos(newPos pixel.Vec) {
	washing.object.SetPos(newPos)
}

func (washing *TakingItInWashingObject) GetState() (WashingState) {
	return washing.state
}

func (washing *TakingItInWashingObject) SetState(newState WashingState) {
	washing.state = newState
}

func (washing *TakingItInWashingObject) Stop() {
	washing.object.HitGround()
}

// Note, this pos is in game screen coordinates
func (washing *TakingItInWashingObject) CheckHit(atPos pixel.Vec) bool {

	checkInMask := atPos.Sub(washing.GetPos()).Scaled(1 / washing.scale).Rotated(-washing.object.GetRotation()).Add(washing.mask.GetRect().Center())
	return washing.mask.GetMask(checkInMask)
}