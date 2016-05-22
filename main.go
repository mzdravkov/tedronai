package main

import (
	"image/color"
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type myScene struct{}

// Type uniquely defines your game type
func (*myScene) Type() string { return "myGame" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*myScene) Preload() {
	engo.Files.Load("textures/city.png")
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (*myScene) Setup(*ecs.World) {
	common.SetBackground(color.White)
	// Retrieve a texture
	texture, err := common.PreloadedSpriteSingle("textures/city.png")
	if err != nil {
		log.Println(err)
	}

}

func main() {
	opts := engo.RunOptions{
		Title:  "Hello World",
		Width:  960,
		Height: 840,
	}
	engo.Run(opts, &myScene{})
}
