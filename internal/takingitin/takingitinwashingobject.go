package takingitin

import (
	"math/rand"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/ext/imdraw"
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
}

func NewWashingObject(
	game *framework.PixelGameController,
	item string,
	startPos pixel.Vec,
	startRotation float64) (*TakingItInWashingObject) {

	scale := (rand.Float64() + 0.2) / 4.0

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
	}
	washing.mask = game.Resources.GetMask(item)
	return &washing
}

func (washing *TakingItInWashingObject) Draw(win *opengl.Window) {

	washing.object.Draw(win)
	
	// DEBUG
	// Check whether the mouse is on this item
	if washing.CheckHit(win.MousePosition()) {
		imd := imdraw.New(nil)
		imd.Color = pixel.RGB(0, 1, 0)
		imd.Push(washing.GetPos())
		imd.Circle(3, 0)
		imd.Draw(win)
	}
}

func (washing *TakingItInWashingObject) HandleKey(win *opengl.Window, dt float64) {
	washing.object.UpdatePos(dt, 20)
}

func (washing *TakingItInWashingObject) Unpeg(horizSpeed float64) {
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

// Note, this pos is in game screen coordinates
func (washing *TakingItInWashingObject) CheckHit(atPos pixel.Vec) bool {

	checkInMask := atPos.Sub(washing.GetPos()).Scaled(1.0 / PEG_SCALE).Rotated(-washing.object.GetRotation()).Add(washing.mask.GetRect().Center())
	return washing.mask.GetMask(checkInMask)
}