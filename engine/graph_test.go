package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGraphAddVertex(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Vertices: map[string]*Vertex{}}

	g.AddVertex("ABC", "Test", WaterTile)

	assert.NotNil(g.Vertices["ABC"])
	assert.Equal("Test", g.Vertices["ABC"].Name)
}

func TestGraphAddEdgeToExistingVertices(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Vertices: map[string]*Vertex{}}

	g.AddVertex("ABC", "Test", WaterTile)
	g.AddVertex("DEF", "Another one", LandTile)
	g.AddEdge("ABC", "DEF")

	assert.NotNil(g.Vertices["ABC"])
	assert.NotNil(g.Vertices["DEF"])
	assert.NotNil(g.Vertices["ABC"].Edges["DEF"])
}

func TestGraphAddEdgeToMissingVertex(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Vertices: map[string]*Vertex{}}

	g.AddVertex("ABC", "Test", WaterTile)
	g.AddEdge("ABC", "DEF")

	assert.NotNil(g.Vertices["ABC"])
	assert.Nil(g.Vertices["DEF"])
	assert.Nil(g.Vertices["ABC"].Edges["DEF"])
}

func TestGetNeighbors(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Vertices: map[string]*Vertex{}}

	g.AddVertex("ABC", "Test 1", WaterTile)
	g.AddVertex("DEF", "Test 2", LandTile)
	g.AddVertex("GHI", "Test 3", WaterTile)
	g.AddVertex("JKL", "Test 4", LandTile)

	g.AddEdge("ABC", "DEF")
	g.AddEdge("ABC", "JKL")
	g.AddEdge("ABC", "GHI")
	g.AddEdge("GHI", "JKL")
	g.AddEdge("GHI", "DEF")

	neighbors := g.GetNeighbors("GHI")
	assert.Equal(2, len(neighbors))
}
