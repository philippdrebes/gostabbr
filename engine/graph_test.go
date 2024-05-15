package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddVertex(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Provinces: map[string]*Province{}}

	g.AddProvince("ABC", "Test", WaterTile, false)

	assert.NotNil(g.Provinces["ABC"])
	assert.Equal("Test", g.Provinces["ABC"].Name)
}

func TestAddEdge_AddEdgeToExistingVertices(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Provinces: map[string]*Province{}}

	g.AddProvince("ABC", "Test", WaterTile, false)
	g.AddProvince("DEF", "Another one", LandTile, false)
	g.AddEdge("ABC", "DEF")

	assert.NotNil(g.Provinces["ABC"])
	assert.NotNil(g.Provinces["DEF"])
	assert.NotNil(g.Provinces["ABC"].Edges["DEF"])
}

func TestAddEdge_AddEdgeToMissingVertex(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Provinces: map[string]*Province{}}

	g.AddProvince("ABC", "Test", WaterTile, false)
	g.AddEdge("ABC", "DEF")

	assert.NotNil(g.Provinces["ABC"])
	assert.Nil(g.Provinces["DEF"])
	assert.Nil(g.Provinces["ABC"].Edges["DEF"])
}

func TestAddEdge_AddEdgesToExistingVertices(t *testing.T) {
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

func TestAddEdge_AddEdgesToMissingVertex(t *testing.T) {
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

	neighbors := g.GetNeighborKeys("GHI")
	assert.Equal(2, len(neighbors))
}

func TestAddUnit_AddUnitToExistingTile(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Provinces: map[string]*Province{}}
	country := &Country{Name: "test country"}
	utype := Fleet

	g.AddProvince("ABC", "Test", WaterTile, false)
	assert.NotNil(g.Provinces["ABC"])

	_, err := g.AddUnit(country, utype, "ABC")

	assert.NoError(err)
	unit := g.Provinces["ABC"].Unit
	assert.NotNil(unit)
	assert.Equal(country, unit.Country)
	assert.Equal(utype, unit.Type)
}

func TestAddUnit_AddUnitToMissingTile(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Provinces: map[string]*Province{}}
	country := &Country{Name: "test country"}
	utype := Fleet

	assert.Nil(g.Provinces["ABC"])

	_, err := g.AddUnit(country, utype, "ABC")

	assert.Error(err)
}

func TestAddUnit_AddUnitToOccupiedTile(t *testing.T) {
	assert := assert.New(t)

	g := &Graph{Provinces: map[string]*Province{}}
	utype := Fleet

	g.AddProvince("ABC", "Test", WaterTile, false)
	assert.NotNil(g.Provinces["ABC"])

	_, err := g.AddUnit(&Country{Name: "country1"}, utype, "ABC")
	assert.NoError(err)

	_, err = g.AddUnit(&Country{Name: "country2"}, utype, "ABC")
	assert.Error(err)
}

func TestGetUnits_NoUnits(t *testing.T) {
	graph := Graph{Provinces: map[string]*Province{}}
	assert.Empty(t, graph.GetUnits("France"), "There should be no units for an empty graph.")
}

func TestGetUnits_UnitsFromMultipleCountries(t *testing.T) {
	france := &Country{Name: "France"}
	germany := &Country{Name: "Germany"}

	graph := Graph{
		Provinces: map[string]*Province{
			"PAR": {Key: "PAR", Unit: &Unit{Country: france, Type: Army}},
			"BER": {Key: "BER", Unit: &Unit{Country: germany, Type: Army}},
			"BRE": {Key: "BRE", Unit: &Unit{Country: france, Type: Fleet}},
		},
	}
	units := graph.GetUnits("France")
	assert.Len(t, units, 2, "There should be two units from France.")
	assert.Contains(t, units, &Unit{Country: france, Type: Army}, "The units should include the army in Paris (PAR).")
	assert.Contains(t, units, &Unit{Country: france, Type: Fleet}, "The units should include the fleet in Brest (BRE).")
}

func TestGetUnits_MultipleUnitsSameCountry(t *testing.T) {
	italy := &Country{Name: "Italy"}

	graph := Graph{
		Provinces: map[string]*Province{
			"ROM": {Key: "ROM"},
			"NAP": {Key: "NAP"},
		},
	}

	rom, _ := graph.AddUnit(italy, Army, "ROM")
	nap, _ := graph.AddUnit(italy, Fleet, "NAP")

	units := graph.GetUnits("Italy")
	assert.Len(t, units, 2, "There should be two units from the same country (Italy).")
	assert.Contains(t, units, rom, "The units should include the army in Rome (ROM).")
	assert.Contains(t, units, nap, "The units should include the fleet in Naples (NAP).")
}

func TestGetUnits_NoneFromSpecifiedCountry(t *testing.T) {
	graph := Graph{
		Provinces: map[string]*Province{
			"MUN": {Key: "MUN", Unit: &Unit{Country: &Country{Name: "Germany"}, Type: Army}},
			"KIE": {Key: "KIE", Unit: &Unit{Country: &Country{Name: "Germany"}, Type: Fleet}},
		},
	}
	assert.Empty(t, graph.GetUnits("Russia"), "There should be no units from the specified country (Russia).")
}

func TestGetProvince_Found(t *testing.T) {
	graph := Graph{
		Provinces: map[string]*Province{
			"PAR": {Key: "PAR", Name: "Paris"},
		},
	}
	province, err := graph.GetProvince("PAR")
	assert.NoError(t, err, "Should not return an error for existing province key.")
	assert.NotNil(t, province, "Returned province should not be nil.")
	assert.Equal(t, "Paris", province.Name, "Province name should match the expected name.")
}

func TestGetProvince_NotFound(t *testing.T) {
	graph := Graph{Provinces: map[string]*Province{}}
	province, err := graph.GetProvince("PAR")
	assert.Error(t, err, "Should return an error for non-existing province key.")
	assert.Nil(t, province, "Returned province should be nil when key does not exist.")
	assert.EqualError(t, err, "Province 'PAR' not found", "Error message should indicate that the province is not found.")
}
