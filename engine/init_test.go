package engine

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

func TestVisualizeInitializedWorld(t *testing.T) {
	assert := assert.New(t)
	filename := "./world.dot"

	game := InitializeNewGame()
	world := game.World
	g := graph.New(graph.StringHash)

	for srcKey, vertex := range world.Provinces {
		addVertex(g, srcKey, vertex)
		for destKey, _ := range vertex.Edges {
			err := g.AddEdge(srcKey, destKey)
			if err != nil {
				addVertex(g, srcKey, vertex)
				_ = g.AddEdge(srcKey, destKey)
			}
		}
	}

	file, err := os.Create(filename)
	assert.NoError(err)
	assert.NoError(draw.DOT(g, file))
	assert.FileExists(filename)
}

func addVertex(g graph.Graph[string, string], srcKey string, vertex *Province) {
	colorscheme := "purples3"
	if vertex.IsSupplyCenter {
		colorscheme = "greens3"
	}
	if vertex.Type == WaterTile {
		colorscheme = "blues3"
	}
	_ = g.AddVertex(srcKey,
		graph.VertexAttribute("colorscheme", colorscheme), graph.VertexAttribute("style", "filled"),
		graph.VertexAttribute("color", "2"), graph.VertexAttribute("fillcolor", "1"))
}
