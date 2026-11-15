package takingitin

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/imdraw"
	"github.com/derekwaters/takingitin/internal/framework"
)

const DEFAULT_BUCKET_SCALE = 0.3

type TakingItInBucketObject struct {
	backSprite	*pixel.Sprite		
	frontSprite	*pixel.Sprite
	pos			pixel.Vec
	mask		*framework.PixelPictureMask
}

func NewTakingItInBucketObject(
	game *framework.PixelGameController,
	backSprite string,
	frontSprite string, 
	startPos pixel.Vec,) (*TakingItInBucketObject) {

	bucket := TakingItInBucketObject{
		pos: startPos,
	}
	backPic := game.Resources.GetPicture(backSprite)
	bucket.backSprite = pixel.NewSprite(backPic, backPic.Bounds())
	frontPic := game.Resources.GetPicture(frontSprite)
	bucket.mask = game.Resources.GetMask(frontSprite)
	bucket.frontSprite = pixel.NewSprite(frontPic, frontPic.Bounds())
	return &bucket
}

func (bucket *TakingItInBucketObject) DrawBack(win *opengl.Window) {
	mat := pixel.IM.Scaled(pixel.ZV, DEFAULT_BUCKET_SCALE).Moved(bucket.pos)
	bucket.backSprite.Draw(win, mat)

	// DEBUG
	// Check whether the mouse is on this item
	imd := imdraw.New(nil)
	if bucket.CheckInBucket(win.MousePosition()) {		
		imd.Color = pixel.RGB(0, 1, 1)
	} else {
		imd.Color = pixel.RGB(1, 1, 0)
	}
	imd.Push(bucket.pos)
	imd.Circle(5, 0)
	imd.Draw(win)
}

func (bucket *TakingItInBucketObject) DrawFront(win *opengl.Window) {
	mat := pixel.IM.Scaled(pixel.ZV, DEFAULT_BUCKET_SCALE).Moved(bucket.pos)
	bucket.frontSprite.Draw(win, mat)
}

func (bucket *TakingItInBucketObject) CheckInBucket(pos pixel.Vec) bool {
	// Transform pos into the bucket image coords
	checkInMask := pos.Sub(bucket.pos).Scaled(1.0 / DEFAULT_BUCKET_SCALE).Add(bucket.mask.GetRect().Center())
	return bucket.mask.GetMask(checkInMask)
}