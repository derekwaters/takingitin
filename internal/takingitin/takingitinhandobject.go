package takingitin

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/derekwaters/takingitin/internal/framework"
)

const MAX_SPEED = 4.0
const ACCEL = 0.5
const FRICTION = 0.3
const DEFAULT_SCALE = 0.4

type TakingItInHandObject struct {
	sprites		[]*pixel.Sprite
	sprite		int
	pos			pixel.Matrix
	speed		pixel.Vec

	upBtn		pixel.Button
	downBtn		pixel.Button
	rightBtn	pixel.Button
	leftBtn		pixel.Button
	closeBtn	pixel.Button
}

func NewTakingItInHandObject(
	game *framework.PixelGameController,
	spriteNames []string, 
	startPos pixel.Vec, 
	upBtn pixel.Button,
	downBtn pixel.Button,
	rightBtn pixel.Button,
	leftBtn pixel.Button,
	closeBtn pixel.Button) (*TakingItInHandObject) {

	hand := TakingItInHandObject{
		pos: pixel.IM.Scaled(pixel.V(0, 0), DEFAULT_SCALE).Moved(startPos),
		speed: pixel.V(0, 0),
		upBtn: upBtn,
		downBtn: downBtn,
		rightBtn: rightBtn,
		leftBtn: leftBtn,
		closeBtn: closeBtn,
	}
	for i := 0; i < len(spriteNames); i++ {
		pic := game.Resources.GetPicture(spriteNames[i])
		hand.sprites = append(hand.sprites, pixel.NewSprite(pic, pic.Bounds()))
	}
	hand.sprite = 0
	return &hand
}

func (hand *TakingItInHandObject) Draw(win *opengl.Window) {
	hand.sprites[hand.sprite].Draw(win, hand.pos)
}

func (hand *TakingItInHandObject) HandleKey(win *opengl.Window) {
	if win.Pressed(hand.leftBtn) {
		hand.speed.X -= ACCEL
	} else if win.Pressed(hand.rightBtn) {
		hand.speed.X += ACCEL
	} else {
		if hand.speed.X < 0 {
			hand.speed.X += FRICTION
			if hand.speed.X > 0 {
				hand.speed.X = 0
			}
		} else if hand.speed.X > 0 {
			hand.speed.X -= FRICTION
			if hand.speed.X < 0 {
				hand.speed.X = 0
			}
		}
	}
	if win.Pressed(hand.downBtn) {
		hand.speed.Y -= ACCEL
	} else if win.Pressed(hand.upBtn) {
		hand.speed.Y += ACCEL
	} else {
		if hand.speed.Y < 0 {
			hand.speed.Y += FRICTION
			if hand.speed.Y > 0 {
				hand.speed.Y = 0
			}
		} else if hand.speed.Y > 0 {
			hand.speed.Y -= FRICTION
			if hand.speed.Y < 0 {
				hand.speed.Y = 0
			}
		}
	}
	// Clamp the speed
	if hand.speed.X < (-1.0 * MAX_SPEED) {
		hand.speed.X = -1.0 * MAX_SPEED
	} else if hand.speed.X > MAX_SPEED {
		hand.speed.X = MAX_SPEED
	}
	if hand.speed.Y < (-1.0 * MAX_SPEED) {
		hand.speed.Y = -1.0 * MAX_SPEED
	} else if hand.speed.Y > MAX_SPEED {
		hand.speed.Y = MAX_SPEED
	}

	if win.Pressed(hand.closeBtn) {
		hand.sprite++
	} else {
		hand.sprite--
	}
	if hand.sprite >= len(hand.sprites) {
		hand.sprite = len(hand.sprites) - 1
	} else if hand.sprite < 0 {
		hand.sprite = 0
	}
	// Now move the hand
	hand.pos = hand.pos.Moved(hand.speed)
}