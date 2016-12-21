package scenes

import (
	"engo.io/engo"
	"engo.io/ecs"
	"engo.io/engo/common"
	"image/color"
	"github.com/tjheslin1/SurvivalGame/systems"
)

type GameScene struct {

}

func (*GameScene) Type() string {
	return "myGame"
}

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*GameScene) Preload() {
	engo.Files.Load("textures/city.png")
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (*GameScene) Setup(world *ecs.World) {
	engo.Input.RegisterAxis("horizontal", engo.AxisKeyPair{engo.D, engo.A})
	engo.Input.RegisterAxis("vertical", engo.AxisKeyPair{engo.S, engo.W})

	common.SetBackground(color.White)

	world.AddSystem(&common.RenderSystem{})
	//world.AddSystem(&common.MouseSystem{})

	kbs := common.NewKeyboardScroller(
		400, engo.DefaultHorizontalAxis,
		engo.DefaultVerticalAxis)
	world.AddSystem(kbs)

	world.AddSystem(&systems.CharacterSystem{})
}