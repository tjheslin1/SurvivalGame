package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"github.com/tjheslin1/SurvivalGame/entities"
	"engo.io/engo/common"
	"fmt"
)

type CharacterSystem struct {
	world     *ecs.World
	character *entities.Character
}

// @Override
func (pm *CharacterSystem) Remove(ecs.BasicEntity) {
	// TODO
}

// @Override
func (pm *CharacterSystem) Update(dt float32) {
	if v := engo.Input.Axis("vertical").Value(); v != 0 {
		pm.character.Position.X += v
	}

	if v := engo.Input.Axis("horizontal").Value(); v != 0 {
		pm.character.Position.Y += v
	}

	fmt.Printf("X: %f      -       Y: %f\n", &pm.character.Position.X, pm.character.Position.Y)

	//if btn := engo.Input.Button("action"); btn.JustPressed() {
	//	fmt.Println("DOWN!")
	//}
}

func (pm *CharacterSystem) New(world *ecs.World) {
	pm.world = world

	pm.character = &entities.Character{BasicEntity: ecs.NewBasic()}

	pm.character.SpaceComponent = common.SpaceComponent{}

	texture, err := common.LoadedSprite("textures/city.png")
	if err != nil {
		panic("Unable to load texture: " + err.Error())
	}

	pm.character.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{X: 0.1, Y: 0.1},
	}

	for _, system := range pm.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&pm.character.BasicEntity, &pm.character.RenderComponent, &pm.character.SpaceComponent)
		}
	}
}