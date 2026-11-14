package takingitin

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/derekwaters/takingitin/internal/framework"
)

const GRAVITY = 9.8


type PhysicsObjectState int

const (
	PhysicsObjectState_Stopped = iota
	PhysicsObjectState_Falling
)


type TakingItInPhysicsObject struct {
	sprites		[]*pixel.Sprite
	sprite		int
	pos			pixel.Matrix
	speed		pixel.Vec
	state		PhysicsObjectState

	rotationSpeed	float64

	defaultScale	float64
}

func NewTakingItInPhysicsObject(
	game *framework.PixelGameController,
	spriteNames []string, 
	scale float64,
	startPos pixel.Vec,
	startSpeed pixel.Vec,
	startRotationSpeed float64) (*TakingItInPhysicsObject) {

	obj := TakingItInPhysicsObject{
		pos: pixel.IM.Scaled(pixel.V(0, 0), scale).Moved(startPos),
		speed: startSpeed,
		state: PhysicsObjectState_Stopped,
		defaultScale: scale,
		rotationSpeed: startRotationSpeed,
	}
	for i := 0; i < len(spriteNames); i++ {
		pic := game.Resources.GetPicture(spriteNames[i])
		obj.sprites = append(obj.sprites, pixel.NewSprite(pic, pic.Bounds()))
	}
	obj.sprite = 0
	return &obj
}

func (obj *TakingItInPhysicsObject) Draw(win *opengl.Window) {
	obj.sprites[obj.sprite].Draw(win, obj.pos)
}

func (obj *TakingItInPhysicsObject) UpdatePos(dt float64) {
	if (obj.state == PhysicsObjectState_Falling) {
		obj.speed.Y -= (GRAVITY * dt)
		obj.pos = obj.pos.Moved(obj.speed).Rotated(pixel.ZV, dt * obj.rotationSpeed)
	}
}

func (obj *TakingItInPhysicsObject) Stop() {
	obj.state = PhysicsObjectState_Stopped
	obj.speed = pixel.V(0, 0)
}

func (obj *TakingItInPhysicsObject) Fall() {
	obj.state = PhysicsObjectState_Falling
}

func (obj *TakingItInPhysicsObject) SetSpeed(speed pixel.Vec) {
	obj.state = PhysicsObjectState_Falling
	obj.speed = speed
}

func (obj *TakingItInPhysicsObject) SetPos(pos pixel.Vec) {
	obj.pos = pixel.IM.Scaled(pixel.V(0, 0), obj.defaultScale).Moved(pos)
}
