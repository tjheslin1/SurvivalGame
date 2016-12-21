package entities

import (
	"engo.io/ecs"
	"engo.io/engo/common"
	"github.com/tjheslin1/SurvivalGame/components"
)

type Character struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	components.SpeedComponent
}

func (*Character) Type() string {
	return "Player"
}
