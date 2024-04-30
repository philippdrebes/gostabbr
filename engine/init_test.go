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

	world := InitializeWorld()
	g := graph.New(graph.StringHash)

	for srcKey, vertex := range world.Vertices {
		_ = g.AddVertex(srcKey)
		for destKey, _ := range vertex.Edges {
			err := g.AddEdge(srcKey, destKey)
			if err != nil {
				_ = g.AddVertex(destKey)
				_ = g.AddEdge(srcKey, destKey)
			}
		}
	}

	file, err := os.Create(filename)
	assert.NoError(err)
	assert.NoError(draw.DOT(g, file))
	assert.FileExists(filename)
}
