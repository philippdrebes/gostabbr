package main

import "gostabbr/engine"

func main() {
	game, err := engine.InitializeNewGame()
	if err != nil {
		panic("this should never happen")
	}

	world := game.World
	for t := range world.Provinces {
		println(world.Provinces[t].Name)
	}
}
