package main

import(
	"fmt"
	"github.com/derekwaters/takingitin/internal/takingitin"
	"github.com/derekwaters/takingitin/internal/framework"
)

func main() {
	fmt.Println("--- Starting The Game ---")

	var theData takingitin.TakingItInGameData
	var theGame framework.PixelGameController
	fmt.Println("--- Launching The Game ---")
	theGame.DebugMode = true
	theGame.Launch(theData)

	fmt.Println("---  Ending The Game  ---")
}