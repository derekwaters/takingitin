package framework

import (
	// "fmt"
	"math"
	"github.com/gopxl/pixel/v2"
)

type PixelPictureMask struct {
	data	[]uint32
	rect	pixel.Rect
	stride	int
}

func NewMaskFromPicture(data *pixel.PictureData) (*PixelPictureMask) {
	w := int(math.Ceil(data.Rect.Max.X)) - int(math.Floor(data.Rect.Min.X))
	h := int(math.Ceil(data.Rect.Max.Y)) - int(math.Floor(data.Rect.Min.Y))
	uintsNeeded := int(math.Ceil(float64(w) * float64(h) / 32.0))

	new := PixelPictureMask{
		rect: data.Rect,
		stride: data.Stride,
	}
	new.data = make([]uint32, uintsNeeded, uintsNeeded)

	for x := 0; x < w; x++ {
		for y:= 0; y < h; y++ {
			x_off := int(math.Floor(data.Rect.Min.X)) + x
			y_off := int(math.Floor(data.Rect.Min.Y)) + y
			isSet := true
			if data.Pix[x_off + (data.Stride * y_off)].A == 0.0 {
				isSet = false
			}
			new.SetMask(pixel.V(float64(x), float64(y)), isSet)
		}
	}
	return &new
}

func (mask *PixelPictureMask) GetMask(atPos pixel.Vec) bool {
	if atPos.X < mask.rect.Min.X || atPos.X > mask.rect.Max.X ||
		atPos.Y < mask.rect.Min.Y || atPos.Y > mask.rect.Max.Y {
		return false
	}
	index, bitMask := mask.GetMaskOffsets(atPos)
	checkVal := mask.data[index] & bitMask
	return checkVal != 0
}

func (mask *PixelPictureMask) SetMask(atPos pixel.Vec, set bool) {
	index, bitMask := mask.GetMaskOffsets(atPos)
	if set {
		mask.data[index] = mask.data[index] | bitMask
	} else {
		mask.data[index] = mask.data[index] & (^bitMask)
	}
}

func (mask *PixelPictureMask) GetRect() (pixel.Rect) {
	return mask.rect
}

func (mask *PixelPictureMask) GetMaskOffsets(atPos pixel.Vec) (int, uint32) {
	y_off := int(math.Floor(atPos.Y - mask.rect.Min.Y))
	x_off := int(math.Floor(atPos.X - mask.rect.Min.X))
	bitPos := (y_off * mask.stride) + x_off


	// 32 bits in each mask, so...
	index := int(math.Floor(float64(bitPos) / 32.0))
	bitOffset := bitPos % 32
	var bitMask uint32

	// fmt.Println(x_off, y_off, bitPos, bitOffset)

	bitMask = 1 << bitOffset
	
	return index, bitMask
}