package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGraphAddVertex(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Provinces: map[string]*Province{}}

	g.AddProvince("ABC", "Test", WaterTile, false)

	assert.NotNil(g.Provinces["ABC"])
	assert.Equal("Test", g.Provinces["ABC"].Name)
}

func TestGraphAddEdgeToExistingVertices(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Provinces: map[string]*Province{}}

	g.AddProvince("ABC", "Test", WaterTile, false)
	g.AddProvince("DEF", "Another one", LandTile, false)
	g.AddEdge("ABC", "DEF")

	assert.NotNil(g.Provinces["ABC"])
	assert.NotNil(g.Provinces["DEF"])
	assert.NotNil(g.Provinces["ABC"].Edges["DEF"])
}

func TestGraphAddEdgeToMissingVertex(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Provinces: map[string]*Province{}}

	g.AddProvince("ABC", "Test", WaterTile, false)
	g.AddEdge("ABC", "DEF")

	assert.NotNil(g.Provinces["ABC"])
	assert.Nil(g.Provinces["DEF"])
	assert.Nil(g.Provinces["ABC"].Edges["DEF"])
}

func TestGraphAddEdgesToExistingVertices(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Provinces: map[string]*Province{}}

	g.AddProvince("ABC", "Test", WaterTile, false)
	g.AddProvince("DEF", "Another one", LandTile, false)
	g.AddProvince("GHI", "Another one", LandTile, false)
	g.AddEdges("ABC", []string{"DEF", "GHI"})

	assert.NotNil(g.Provinces["ABC"])
	assert.NotNil(g.Provinces["DEF"])
	assert.NotNil(g.Provinces["ABC"].Edges["DEF"])
	assert.NotNil(g.Provinces["ABC"].Edges["GHI"])
}

func TestGraphAddEdgesToMissingVertex(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Provinces: map[string]*Province{}}

	g.AddProvince("ABC", "Test", WaterTile, false)
	g.AddEdges("ABC", []string{"DEF", "GHI"})

	assert.NotNil(g.Provinces["ABC"])
	assert.Nil(g.Provinces["DEF"])
	assert.Nil(g.Provinces["ABC"].Edges["DEF"])
	assert.Nil(g.Provinces["ABC"].Edges["GHI"])
}

func TestGetNeighbors(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Provinces: map[string]*Province{}}

	g.AddProvince("ABC", "Test 1", WaterTile, false)
	g.AddProvince("DEF", "Test 2", LandTile, false)
	g.AddProvince("GHI", "Test 3", WaterTile, false)
	g.AddProvince("JKL", "Test 4", LandTile, false)

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

	g := &Graph{Provinces: map[string]*Province{}}
	country := "test country"
	utype := Fleet

	g.AddProvince("ABC", "Test", WaterTile, false)
	assert.NotNil(g.Provinces["ABC"])

	err := g.AddUnit(country, utype, "ABC")

	assert.NoError(err)
	unit := g.Provinces["ABC"].Unit
	assert.NotNil(unit)
	assert.Equal(country, unit.Country)
	assert.Equal(utype, unit.Type)
}

func TestGraphAddUnitToMissingTile(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Provinces: map[string]*Province{}}
	country := "test country"
	utype := Fleet

	assert.Nil(g.Provinces["ABC"])

	err := g.AddUnit(country, utype, "ABC")

	assert.Error(err)
}
