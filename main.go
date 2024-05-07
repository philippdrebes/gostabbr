package main

import "gostabbr/engine"

func main() {
	game := engine.InitializeNewGame()
	world := game.World
	for t := range world.Provinces {
		println(world.Provinces[t].Name)
	}
}
