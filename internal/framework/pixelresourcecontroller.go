package framework

import (
	"os"
	"image"
	"io/ioutil"
	"image/color"

	_ "image/jpeg"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/ext/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type PixelResourceController struct {
	fonts		map[string]*text.Atlas;
	pictures	map[string]pixel.Picture;
}

func (controller *PixelResourceController) Init(game *PixelGameController) {
	controller.fonts = make(map[string]*text.Atlas)
	controller.pictures = make(map[string]pixel.Picture)
}

func (controller *PixelResourceController) GenerateTextElement(pos pixel.Vec, tag string, color color.RGBA) *text.Text {
	item := text.New(pos, controller.GetFont(tag))
	item.Color = color
	return item
}

func (controller *PixelResourceController) GetPicture(tag string) (pixel.Picture) {
	pix, _ := controller.pictures[tag]
	return pix
}

func (controller *PixelResourceController) GetFont(tag string) (*text.Atlas) {
	return controller.fonts[tag]
}

func (controller *PixelResourceController) AddPicture(tag string, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	controller.pictures[tag] = pixel.PictureDataFromImage(img)
	return nil
}

func (controller *PixelResourceController) AddFont(tag string, path string, size float64) error {
	face, err := controller.loadTTF(path, size)
	if err != nil {
		return err
	}
	controller.fonts[tag] = text.NewAtlas(face, text.ASCII)
	return nil
}

func (controller *PixelResourceController) loadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}

