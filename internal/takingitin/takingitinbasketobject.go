package takingitin

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	//"github.com/gopxl/pixel/v2/ext/imdraw"
	"github.com/derekwaters/takingitin/internal/framework"
)

const DEFAULT_BASKET_SCALE = 0.4

type TakingItInBasketObject struct {
	backSprite	*pixel.Sprite		
	frontSprite	*pixel.Sprite
	pos			pixel.Vec
	extent		pixel.Rect
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
	basket.extent = backPic.Bounds()
	frontPic := game.Resources.GetPicture(frontSprite)
	basket.frontSprite = pixel.NewSprite(frontPic, frontPic.Bounds())
	return &basket
}

func (basket *TakingItInBasketObject) DrawBack(win *opengl.Window) {
	mat := pixel.IM.Scaled(pixel.ZV, DEFAULT_BASKET_SCALE).Moved(basket.pos)
	basket.backSprite.Draw(win, mat)
}

func (basket *TakingItInBasketObject) DrawFront(win *opengl.Window) {
	mat := pixel.IM.Scaled(pixel.ZV, DEFAULT_BASKET_SCALE).Moved(basket.pos)
	basket.frontSprite.Draw(win, mat)
}

func (basket *TakingItInBasketObject) CheckInBasket(pos pixel.Vec) bool {
	// Transform pos into the basket image coords
	checkVec := pos.Sub(basket.pos).Scaled(DEFAULT_BASKET_SCALE)
	return basket.extent.Contains(checkVec)
}