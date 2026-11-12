package framework

import (
	"github.com/gopxl/pixel/v2"
)

type PixelGameData interface {
	GetTitle() string
	GetBounds() pixel.Rect
	Init(controller *PixelGameController)
}
