package engine

import "errors"

type TileType int8
type UnitType int8

const (
	LandTile TileType = iota
	WaterTile
)

const (
	Army UnitType = iota
	Fleet
)

type Graph struct {
	Provinces map[string]*Province
}

type Province struct {
	Key            string
	Name           string
	Type           TileType
	IsSupplyCenter bool
	OwnedBy        string
	Unit           *Unit
	Edges          map[string]*Edge
}

type Edge struct {
	Vertex *Province
}

type Unit struct {
	Country string
	Type    UnitType
}

func (this *Graph) AddProvince(key, name string, tileType TileType, isSupplyCenter bool) {
	this.Provinces[key] = &Province{Key: key, Name: name, Type: tileType, IsSupplyCenter: isSupplyCenter, Edges: map[string]*Edge{}}
}

func (this *Graph) AddEdge(srcKey, destKey string) {
	if _, ok := this.Provinces[srcKey]; !ok {
		return
	}
	if _, ok := this.Provinces[destKey]; !ok {
		return
	}

	this.Provinces[srcKey].Edges[destKey] = &Edge{Vertex: this.Provinces[destKey]}
}

func (this *Graph) AddUnit(country string, unitType UnitType, tile string) error {
	if _, ok := this.Provinces[tile]; !ok {
		return errors.New("Tile not found")
	}

	unit := Unit{
		Country: country,
		Type:    unitType,
	}

	this.Provinces[tile].Unit = &unit

	return nil
}

func (this *Graph) AddEdges(srcKey string, destKeys []string) {
	if _, ok := this.Provinces[srcKey]; !ok {
		return
	}

	for _, destKey := range destKeys {
		if _, ok := this.Provinces[destKey]; !ok {
			return
		}

		this.Provinces[srcKey].Edges[destKey] = &Edge{Vertex: this.Provinces[destKey]}
	}
}

func (this *Graph) GetNeighbors(srcKey string) []string {
	result := []string{}

	for _, edge := range this.Provinces[srcKey].Edges {
		result = append(result, edge.Vertex.Name)
	}

	return result
}
