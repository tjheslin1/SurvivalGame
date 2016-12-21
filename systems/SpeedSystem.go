package systems

import (
	"engo.io/ecs"
	"engo.io/engo/common"
	"engo.io/engo"
	"log"
	"github.com/tjheslin1/SurvivalGame/components"
)

const (
	SPEED_MESSAGE = "SpeedMessage"
	SPEED_SCALE = 24
)

type SpeedMessage struct {
	*ecs.BasicEntity
	engo.Point
}

func (SpeedMessage) Type() string {
	return SPEED_MESSAGE
}

type speedEntity struct {
	*ecs.BasicEntity
	*components.SpeedComponent
	*common.SpaceComponent
}

type SpeedSystem struct {
	entities []speedEntity
}

func (s *SpeedSystem) New(*ecs.World) {
	engo.Mailbox.Listen(SPEED_MESSAGE, func(message engo.Message) {
		speed, isSpeed := message.(SpeedMessage)
		if isSpeed {
			log.Printf("%#v\n", speed.Point)
			for _, e := range s.entities {
				if e.ID() == speed.BasicEntity.ID() {
					e.SpeedComponent.Point = speed.Point
				}
			}
		}
	})
}

func (s *SpeedSystem) Add(basic *ecs.BasicEntity, speed *components.SpeedComponent, space *common.SpaceComponent) {
	s.entities = append(s.entities, speedEntity{basic, speed, space})
}

func (s *SpeedSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range s.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		s.entities = append(s.entities[:delete], s.entities[delete + 1:]...)
	}
}

func (s *SpeedSystem) Update(dt float32) {
	for _, entity := range s.entities {
		speed := engo.GameWidth() * dt
		entity.SpaceComponent.Position.X = entity.SpaceComponent.Position.X + speed * entity.SpeedComponent.Point.X
		entity.SpaceComponent.Position.Y = entity.SpaceComponent.Position.Y + speed * entity.SpeedComponent.Point.Y

		// Add Game Border Limits
		// TODO calculate border limits
		var heightLimit float32 = 400 - entity.SpaceComponent.Height
		if entity.SpaceComponent.Position.Y < 0 {
			entity.SpaceComponent.Position.Y = 0
		} else if entity.SpaceComponent.Position.Y > heightLimit {
			entity.SpaceComponent.Position.Y = heightLimit
		}

		var widthLimit float32 = 400 - entity.SpaceComponent.Width
		if entity.SpaceComponent.Position.X < 0 {
			entity.SpaceComponent.Position.X = 0
		} else if entity.SpaceComponent.Position.X > widthLimit {
			entity.SpaceComponent.Position.X = widthLimit
		}
	}

}
