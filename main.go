package main

import "gostabbr/engine"

func main() {
	game := engine.InitializeNewGame()
	world := game.World
	for t := range world.Vertices {
		println(world.Vertices[t].Name)
	}
}
