package takingitin

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/derekwaters/takingitin/internal/framework"
)

const DEFAULT_BUCKET_SCALE = 0.3

type TakingItInBucketObject struct {
	backSprite	*pixel.Sprite		
	frontSprite	*pixel.Sprite
	pos			pixel.Vec
	extent		pixel.Rect
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
	bucket.extent = backPic.Bounds()
	frontPic := game.Resources.GetPicture(frontSprite)
	bucket.frontSprite = pixel.NewSprite(frontPic, frontPic.Bounds())
	return &bucket
}

func (bucket *TakingItInBucketObject) DrawBack(win *opengl.Window) {
	mat := pixel.IM.Scaled(pixel.ZV, DEFAULT_BUCKET_SCALE).Moved(bucket.pos)
	bucket.backSprite.Draw(win, mat)
}

func (bucket *TakingItInBucketObject) DrawFront(win *opengl.Window) {
	mat := pixel.IM.Scaled(pixel.ZV, DEFAULT_BUCKET_SCALE).Moved(bucket.pos)
	bucket.frontSprite.Draw(win, mat)
}

func (bucket *TakingItInBucketObject) CheckInBucket(pos pixel.Vec) bool {
	// Transform pos into the bucket image coords
	checkVec := pos.Sub(bucket.pos).Scaled(DEFAULT_BUCKET_SCALE)
	return bucket.extent.Contains(checkVec)
}