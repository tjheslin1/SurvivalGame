package systems

import (
	"fmt"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

	"github.com/tjheslin1/SurvivalGame/io"
	"github.com/tjheslin1/SurvivalGame/entities"
)

type CityBuildingSystem struct {
	world        *ecs.World
	mouseTracker io.MouseTracker
}

// Remove is called whenever an Entity is removed from the World, in order to remove it from this system as well
func (*CityBuildingSystem) Remove(ecs.BasicEntity) {

}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (cb *CityBuildingSystem) Update(dt float32) {
	if engo.Input.Button("AddCity").JustPressed() {
		fmt.Println("The gamer pressed 'C'")

		city := entities.City{BasicEntity: ecs.NewBasic()}

		city.SpaceComponent = common.SpaceComponent{
			Position: engo.Point{cb.mouseTracker.MouseX, cb.mouseTracker.MouseY},
			Width:    30,
			Height:   64,
		}

		texture, err := common.LoadedSprite("textures/city.png")
		if err != nil {
			panic("Unable to load texture: " + err.Error())
		}

		city.RenderComponent = common.RenderComponent{
			Drawable: texture,
			Scale:    engo.Point{X: 0.1, Y: 0.1},
		}

		for _, system := range cb.world.Systems() {
			switch sys := system.(type) {
			case *common.RenderSystem:
				sys.Add(&city.BasicEntity, &city.RenderComponent, &city.SpaceComponent)
			}
		}
	}
}

// New is the initialisation of the System
func (cb *CityBuildingSystem) New(world *ecs.World) {
	cb.world = world

	fmt.Println("CityBuildingSystem was added to the Scene")

	cb.mouseTracker.BasicEntity = ecs.NewBasic()
	cb.mouseTracker.MouseComponent = common.MouseComponent{Track: true}

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(&cb.mouseTracker.BasicEntity, &cb.mouseTracker.MouseComponent, nil, nil)
		}
	}
}