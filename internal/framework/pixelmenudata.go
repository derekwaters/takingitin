package framework

import (
	"image/color"
)

type PixelMenuSelection int

const (
	PixelMenuSelection_NewGame = iota
	PixelMenuSelection_Options
	PixelMenuSelection_About
	PixelMenuSelection_Quit
)

var PixelMenuSelectionName = map[PixelMenuSelection]string{
	PixelMenuSelection_NewGame:		"New Game",
	PixelMenuSelection_Options:		"Options",
	PixelMenuSelection_About:		"About",
	PixelMenuSelection_Quit:		"Quit",
}

type PixelMenuData struct {
	BackgroundColor	color.RGBA
	MenuTextColor color.RGBA
	BackgroundImageTag	string
	TitleFontTag string
	MenuFontTag	string
	MenuOptions []PixelMenuSelection
	NewGameSceneName string
	AboutSceneName string
}
