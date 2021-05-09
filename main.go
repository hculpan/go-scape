package main

import (
	"embed"
	"fmt"

	"github.com/hculpan/go-scape/app/controllers"
	"github.com/hculpan/go-sdl-lib/component"
	"github.com/hculpan/go-sdl-lib/resources"
)

func main() {
	component.SetupSDL()

	if err := resources.FontsInit(embed.FS{}); err != nil {
		fmt.Println(err)
		return
	}

	// Since our cells are all 3 pixels with a 1 pixel barrier
	// around them, we want to make sure our widht/height is
	// a divisor of 4
	var gameWidth int32 = 809
	var gameHeight int32 = 849

	gamecontroller := controllers.NewSugarscapeController(gameWidth, gameHeight)
	if err := gamecontroller.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
