package main

import "gostabbr/engine"

func main() {
	world := engine.InitializeGraph()
	for t := range world.Vertices {
		println(world.Vertices[t].Name)
	}
}
