package main

import (
	"engo.io/engo"
	"github.com/tjheslin1/SurvivalGame/scenes"
)

func main() {
	opts := engo.RunOptions{
		Title: "Hello World",
		Width:  400,
		Height: 400,
		StandardInputs: true,
	}
	engo.Run(opts, &scenes.GameScene{})
}
