package entities

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type Character struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (*Character) Type() string {
	return "Player"
}

func (*Character) New() {

}
