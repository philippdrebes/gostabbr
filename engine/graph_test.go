package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGraphAddVertex(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Vertices: map[string]*Vertex{}}

	g.AddVertex("ABC", "Test", WaterTile, false)

	assert.NotNil(g.Vertices["ABC"])
	assert.Equal("Test", g.Vertices["ABC"].Name)
}

func TestGraphAddEdgeToExistingVertices(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Vertices: map[string]*Vertex{}}

	g.AddVertex("ABC", "Test", WaterTile, false)
	g.AddVertex("DEF", "Another one", LandTile, false)
	g.AddEdge("ABC", "DEF")

	assert.NotNil(g.Vertices["ABC"])
	assert.NotNil(g.Vertices["DEF"])
	assert.NotNil(g.Vertices["ABC"].Edges["DEF"])
}

func TestGraphAddEdgeToMissingVertex(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Vertices: map[string]*Vertex{}}

	g.AddVertex("ABC", "Test", WaterTile, false)
	g.AddEdge("ABC", "DEF")

	assert.NotNil(g.Vertices["ABC"])
	assert.Nil(g.Vertices["DEF"])
	assert.Nil(g.Vertices["ABC"].Edges["DEF"])
}

func TestGraphAddEdgesToExistingVertices(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Vertices: map[string]*Vertex{}}

	g.AddVertex("ABC", "Test", WaterTile, false)
	g.AddVertex("DEF", "Another one", LandTile, false)
	g.AddVertex("GHI", "Another one", LandTile, false)
	g.AddEdges("ABC", []string{"DEF", "GHI"})

	assert.NotNil(g.Vertices["ABC"])
	assert.NotNil(g.Vertices["DEF"])
	assert.NotNil(g.Vertices["ABC"].Edges["DEF"])
	assert.NotNil(g.Vertices["ABC"].Edges["GHI"])
}

func TestGraphAddEdgesToMissingVertex(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Vertices: map[string]*Vertex{}}

	g.AddVertex("ABC", "Test", WaterTile, false)
	g.AddEdges("ABC", []string{"DEF", "GHI"})

	assert.NotNil(g.Vertices["ABC"])
	assert.Nil(g.Vertices["DEF"])
	assert.Nil(g.Vertices["ABC"].Edges["DEF"])
	assert.Nil(g.Vertices["ABC"].Edges["GHI"])
}

func TestGetNeighbors(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Vertices: map[string]*Vertex{}}

	g.AddVertex("ABC", "Test 1", WaterTile, false)
	g.AddVertex("DEF", "Test 2", LandTile, false)
	g.AddVertex("GHI", "Test 3", WaterTile, false)
	g.AddVertex("JKL", "Test 4", LandTile, false)

	g.AddEdge("ABC", "DEF")
	g.AddEdge("ABC", "JKL")
	g.AddEdge("ABC", "GHI")
	g.AddEdge("GHI", "JKL")
	g.AddEdge("GHI", "DEF")

	neighbors := g.GetNeighbors("GHI")
	assert.Equal(2, len(neighbors))
}

func TestGraphAddUnitToExistingTile(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Vertices: map[string]*Vertex{}}
	country := "test country"
	utype := Fleet

	g.AddVertex("ABC", "Test", WaterTile, false)
	assert.NotNil(g.Vertices["ABC"])

	err := g.AddUnit(country, utype, "ABC")

	assert.NoError(err)
	unit := g.Vertices["ABC"].Unit
	assert.NotNil(unit)
	assert.Equal(country, unit.Country)
	assert.Equal(utype, unit.Type)
}

func TestGraphAddUnitToMissingTile(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Vertices: map[string]*Vertex{}}
	country := "test country"
	utype := Fleet

	assert.Nil(g.Vertices["ABC"])

	err := g.AddUnit(country, utype, "ABC")

	assert.Error(err)
}
