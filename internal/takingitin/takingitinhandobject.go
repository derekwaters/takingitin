package takingitin

import (
	"math"
	"math/rand"
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
//	"github.com/gopxl/pixel/v2/ext/imdraw"
	"github.com/derekwaters/takingitin/internal/framework"
)

const MAX_SPEED = 400.0
const ACCEL = 50
const FRICTION = 30
const DEFAULT_HAND_SCALE = 0.2

const WASHING_GRIP_PRESSURE_INCREASE = 0.5
const WASHING_GRIP_PRESSURE_THRESHOLD = 50.0
const WASHING_GRIP_PRESSURE_VIBRATION = 10.0

type TakingItInHandObject struct {
	sprites		[]*pixel.Sprite
	sprite		int
	pos			pixel.Vec
	speed		pixel.Vec

	upBtn		pixel.Button
	downBtn		pixel.Button
	rightBtn	pixel.Button
	leftBtn		pixel.Button
	closeBtn	pixel.Button

	gripOffset			pixel.Vec
	grippingPeg			*TakingItInPegObject
	grippingWashing		*TakingItInWashingObject
	washingGripPressure	float64
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
		pos: startPos,
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
	mat := pixel.IM.Scaled(pixel.ZV, DEFAULT_HAND_SCALE).Moved(hand.pos)
	hand.sprites[hand.sprite].Draw(win, mat)

	// Debug
	// imd := imdraw.New(nil)
	// if hand.grippingPeg != nil {
	// 	imd.Color = pixel.RGB(0, 1, 0)
	// } else if hand.grippingWashing != nil {
	// 	imd.Color = pixel.RGB(0, 0, 1)
	// } else {
	// 	imd.Color = pixel.RGB(1, 0, 0)
	// }
	// imd.Push(hand.pos)
	// imd.Circle(5, 0)
	// imd.Draw(win)
}

func (hand *TakingItInHandObject) HandleKey(win *opengl.Window, dt float64) {
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

	// If the hand is about to move:
	// Now check the state
	// if we're gripping a peg
	//   peg is on the line?
	//     unpeg it
	//     check the clothes its attached to
	//   else
	//     move the peg
	// else if we're gripping washing
	//   if the washing is on the line
	//     increase drag pressure
	//     if the drag pressure exceeds a value
	//        send the pegs flying
	//     else
	//        wobble position a little
	//   else
	//     move the washing

	addToPos := hand.speed.Scaled(dt)

	if hand.speed.X != 0 || hand.speed.Y != 0 {
		if hand.grippingPeg != nil {
			if hand.grippingPeg.GetState() == PegState_OnTheLine {
				// This will check to see if the clothes are free
				hand.grippingPeg.Unpeg()
			}
		} else if hand.grippingWashing != nil {
			if hand.grippingWashing.GetState() == WashingState_OnTheLine {
				hand.washingGripPressure += WASHING_GRIP_PRESSURE_INCREASE
				if hand.washingGripPressure > WASHING_GRIP_PRESSURE_THRESHOLD {
					// Let it fly!
					hand.grippingWashing.Unpeg()
				} else {
					vibrate := (hand.washingGripPressure / WASHING_GRIP_PRESSURE_THRESHOLD) *
						rand.Float64() * WASHING_GRIP_PRESSURE_VIBRATION
					addToPos = pixel.ZV.Unit().Rotated(rand.Float64() * 2 * math.Pi).Scaled(vibrate)
				}
			}
		}
	}
	// Now move the hand
	hand.pos = hand.pos.Add(addToPos)
}

func (hand *TakingItInHandObject) HandleGripping(scene *TakingItInWashingScene) {
	if hand.sprite == (len(hand.sprites) - 1) {
		// Gripping
		if hand.grippingPeg != nil {
			// We already had a peg, so move it
			hand.grippingPeg.SetPos(hand.pos.Sub(hand.gripOffset))
		} else if hand.grippingWashing != nil {
			// We already had washing, so move it
			hand.grippingWashing.SetPos(hand.pos.Sub(hand.gripOffset))
		} else {
			hand.grippingPeg = scene.FindGrabbedPeg(hand.pos)
			if hand.grippingPeg != nil {
				// We got something!
				hand.gripOffset = hand.pos.Sub(hand.grippingPeg.GetPos())
			} else {
				hand.grippingWashing = scene.FindGrabbedWashing(hand.pos)
				if hand.grippingWashing != nil {
					hand.gripOffset = hand.pos.Sub(hand.grippingWashing.GetPos())
					hand.washingGripPressure = 0
				}
			}
		}
	} else {
		// If we were gripping something, then drop it
		if hand.grippingPeg != nil {
			hand.grippingPeg.Drop(hand.speed.X)
			hand.grippingPeg = nil
			hand.gripOffset = pixel.ZV
		} else if hand.grippingWashing != nil {
			hand.grippingWashing.Drop(hand.speed.X / 3)
			hand.grippingWashing = nil
			hand.gripOffset = pixel.ZV
		}
	}
}