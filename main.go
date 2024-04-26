package main

import "gostabbr/engine"

func main() {
	world := engine.InitializeWorld()
	for t := range world.Vertices {
		println(world.Vertices[t].Name)
	}
}
