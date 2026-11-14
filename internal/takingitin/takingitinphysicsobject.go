package takingitin

import (
	"math"
	"math/rand"
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/derekwaters/takingitin/internal/framework"
)

const GRAVITY = 400
const RESTITUTION = 0.5
const STOPPED_SPEED = 0.1
const ROTATION_SCALE = 10.0


type PhysicsObjectState int
type PhysicsGroundCollisionMode int

const (
	PhysicsObjectState_Stopped = iota
	PhysicsObjectState_Falling
)

const (
	PhysicsGroundCollisionMode_BounceDirect = iota
	PhysicsGroundCollisionMode_BounceRandom
	PhysicsGroundCollisionMode_Stop
)


type TakingItInPhysicsObject struct {
	sprites		[]*pixel.Sprite
	sprite		int

	pos			pixel.Vec
	speed		pixel.Vec
	state		PhysicsObjectState

	rotation		float64
	rotationSpeed	float64

	defaultScale	float64

	groundMode	PhysicsGroundCollisionMode
}

func NewTakingItInPhysicsObject(
	game *framework.PixelGameController,
	spriteNames []string, 
	scale float64,
	startPos pixel.Vec,
	startSpeed pixel.Vec,
	startRotation float64,
	startRotationSpeed float64,
	groundMode PhysicsGroundCollisionMode) (*TakingItInPhysicsObject) {

	obj := TakingItInPhysicsObject{
		pos: startPos,
		speed: startSpeed,
		state: PhysicsObjectState_Stopped,
		defaultScale: scale,
		rotation: startRotation,
		rotationSpeed: startRotationSpeed,
		groundMode: groundMode,
	}
	for i := 0; i < len(spriteNames); i++ {
		pic := game.Resources.GetPicture(spriteNames[i])
		obj.sprites = append(obj.sprites, pixel.NewSprite(pic, pic.Bounds()))
	}
	obj.sprite = 0
	return &obj
}

func (obj *TakingItInPhysicsObject) Draw(win *opengl.Window) {
	mat := pixel.IM.Rotated(pixel.ZV, obj.rotation).Scaled(pixel.ZV, obj.defaultScale).Moved(obj.pos)
	obj.sprites[obj.sprite].Draw(win, mat)
}

func (obj *TakingItInPhysicsObject) UpdatePos(dt float64, groundLevel float64) {
	if (obj.state == PhysicsObjectState_Falling) {
		obj.speed.Y -= (GRAVITY * dt)
		obj.rotation += obj.rotationSpeed * dt * ROTATION_SCALE

		// If the object would have hit the ground, then bounce it
		if (obj.pos.Y + (obj.speed.Y * dt)) < groundLevel {
			if obj.groundMode == PhysicsGroundCollisionMode_BounceDirect {
				obj.speed.Y = -obj.speed.Y
				obj.speed = obj.speed.Scaled(RESTITUTION)
				if obj.speed.Len() < STOPPED_SPEED {
					obj.Stop()
				}
			} else if obj.groundMode == PhysicsGroundCollisionMode_BounceRandom {
				obj.rotationSpeed = 0.1 * (rand.Float64() - 0.5)
				origSpeed := obj.speed.Len()
				obj.speed = pixel.V(1, 0).Rotated(math.Pi * rand.Float64()).Scaled(origSpeed * RESTITUTION)
				if obj.speed.Len() < STOPPED_SPEED {
					obj.Stop()
				}
			} else {
				obj.Stop()
				obj.pos.Y = groundLevel
			}
		} else {
			obj.pos = obj.pos.Add(obj.speed.Scaled(dt))
		}
	}
}

func (obj *TakingItInPhysicsObject) Stop() {
	obj.state = PhysicsObjectState_Stopped
	obj.speed = pixel.ZV
	obj.rotationSpeed = 0
}

func (obj *TakingItInPhysicsObject) Fall(rotationSpeed float64) {
	obj.rotationSpeed = rotationSpeed
	obj.state = PhysicsObjectState_Falling
}

func (obj *TakingItInPhysicsObject) SetSpeed(speed pixel.Vec) {
	obj.state = PhysicsObjectState_Falling
	obj.speed = speed
}

func (obj *TakingItInPhysicsObject) SetPos(pos pixel.Vec) {
	obj.pos = pos
}

func (obj *TakingItInPhysicsObject) GetPos() (pos pixel.Vec) {
	return obj.pos
}

func (obj *TakingItInPhysicsObject) GetRotation() (float64) {
	return obj.rotation
}

func (obj *TakingItInPhysicsObject) GetSprite() (*pixel.Sprite) {
	return obj.sprites[obj.sprite]
}