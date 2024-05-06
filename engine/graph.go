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
	Vertices map[string]*Vertex
}

type Vertex struct {
	Key            string
	Name           string
	Type           TileType
	IsSupplyCenter bool
	OwnedBy        string
	Unit           *Unit
	Edges          map[string]*Edge
}

type Edge struct {
	Vertex *Vertex
}

type Unit struct {
	Country string
	Type    UnitType
}

func (this *Graph) AddVertex(key, name string, tileType TileType, isSupplyCenter bool) {
	this.Vertices[key] = &Vertex{Key: key, Name: name, Type: tileType, IsSupplyCenter: isSupplyCenter, Edges: map[string]*Edge{}}
}

func (this *Graph) AddEdge(srcKey, destKey string) {
	if _, ok := this.Vertices[srcKey]; !ok {
		return
	}
	if _, ok := this.Vertices[destKey]; !ok {
		return
	}

	this.Vertices[srcKey].Edges[destKey] = &Edge{Vertex: this.Vertices[destKey]}
}

func (this *Graph) AddUnit(country string, unitType UnitType, tile string) error {
	if _, ok := this.Vertices[tile]; !ok {
		return errors.New("Tile not found")
	}

	unit := Unit{
		Country: country,
		Type:    unitType,
	}

	this.Vertices[tile].Unit = &unit

	return nil
}

func (this *Graph) AddEdges(srcKey string, destKeys []string) {
	if _, ok := this.Vertices[srcKey]; !ok {
		return
	}

	for _, destKey := range destKeys {
		if _, ok := this.Vertices[destKey]; !ok {
			return
		}

		this.Vertices[srcKey].Edges[destKey] = &Edge{Vertex: this.Vertices[destKey]}
	}
}

func (this *Graph) GetNeighbors(srcKey string) []string {
	result := []string{}

	for _, edge := range this.Vertices[srcKey].Edges {
		result = append(result, edge.Vertex.Name)
	}

	return result
}
