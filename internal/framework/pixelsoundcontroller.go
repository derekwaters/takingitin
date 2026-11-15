package framework

import (
	"fmt"
	// "path/filepath"
)

// TODO: 

type PixelSoundController struct {
	// sounds		map[string]*wavPlayer.WavPlayer
}

func (controller *PixelSoundController) Init(game *PixelGameController) {
	// controller.sounds = make(map[string]*wavPlayer.WavPlayer)
}

func (controller *PixelSoundController) AddSounds(defs map[string]string) error {
	for k, v := range(defs) {
		err := controller.AddSound(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (controller *PixelSoundController) AddSound(tag string, path string) error {
	fmt.Println("Loading", path)
	/*
	player, err := wavPlayer.NewWAV(path)
	if err != nil {
		fmt.Println("Got an err", err)
		return err
	}
	// controller.sounds[tag] = player
	*/
	return nil
}

func (controller *PixelSoundController) StopAll() {
}

func (controller *PixelSoundController) Play(tag string) {
}

func (controller *PixelSoundController) Loop(tag string) {
}
