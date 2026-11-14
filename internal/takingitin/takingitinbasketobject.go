package takingitin

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	//"github.com/gopxl/pixel/v2/ext/imdraw"
	"github.com/derekwaters/takingitin/internal/framework"
)

type TakingItInBasketObject struct {
	backSprite	*pixel.Sprite		
	frontSprite	*pixel.Sprite
	pos			pixel.Vec
}

func NewTakingItInBasketObject(
	game *framework.PixelGameController,
	backSprite string,
	frontSprite string, 
	startPos pixel.Vec,) (*TakingItInBasketObject) {

	basket := TakingItInBasketObject{
		pos: startPos,
	}
	backPic := game.Resources.GetPicture(backSprite)
	basket.backSprite = pixel.NewSprite(backPic, backPic.Bounds())
	frontPic := game.Resources.GetPicture(frontSprite)
	basket.frontSprite = pixel.NewSprite(frontPic, frontPic.Bounds())
	return &basket
}

func (basket *TakingItInBasketObject) DrawBack(win *opengl.Window) {
	mat := pixel.IM.Scaled(pixel.ZV, DEFAULT_SCALE).Moved(basket.pos)
	basket.backSprite.Draw(win, mat)
}

func (basket *TakingItInBasketObject) DrawFront(win *opengl.Window) {
	mat := pixel.IM.Scaled(pixel.ZV, DEFAULT_SCALE).Moved(basket.pos)
	basket.frontSprite.Draw(win, mat)
}
