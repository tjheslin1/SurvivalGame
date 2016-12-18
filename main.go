package main

import (
	"engo.io/engo"
	"engo.io/engo/common"
	"engo.io/ecs"
	"image/color"

	"github.com/tjheslin1/SurvivalGame/systems"
)

type myScene struct {

}

func (*myScene) Type() string {
	return "myGame"
}

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*myScene) Preload() {
	engo.Files.Load("textures/city.png")
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (*myScene) Setup(world *ecs.World) {
	cityBuildingSystem := &systems.CityBuildingSystem{}
	engo.Input.RegisterButton("AddCity", engo.C)

	common.SetBackground(color.White)

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})

	world.AddSystem(cityBuildingSystem)
}

func main() {
	opts := engo.RunOptions{
		Title: "Hello World",
		Width:  400,
		Height: 400,
	}
	engo.Run(opts, &myScene{})
}
