package framework

import (
	"fmt"
	"os"
	"log"
	// "path/filepath"
	"time"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)


type PixelSoundController struct {
	sounds				map[string]*beep.Buffer
	speakerInitialised	bool
}

func (controller *PixelSoundController) Init(game *PixelGameController) {
	controller.sounds = make(map[string]*beep.Buffer)
	controller.speakerInitialised = false
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

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()

	if !controller.speakerInitialised {
		fmt.Println("speaker inited")
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		controller.speakerInitialised = true
	}
	controller.sounds[tag] = buffer

	return nil
}

func (controller *PixelSoundController) StopAll() {
	fmt.Println("CLEARING")
	speaker.Clear()
}

func (controller *PixelSoundController) Play(tag string) {
	fmt.Println("HERE WE GO")
	stream := controller.sounds[tag].Streamer(0, controller.sounds[tag].Len())
	speaker.Play(stream)
}

func (controller *PixelSoundController) Loop(tag string) {
	fmt.Println("HERE WE GO 2")
	stream := controller.sounds[tag].Streamer(0, controller.sounds[tag].Len())
	speaker.Play(stream)
}
