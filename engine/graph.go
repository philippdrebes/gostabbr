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
	Province *Province
}

type Unit struct {
	Country string
	Type    UnitType
}

func (this *Graph) AddProvince(key, name string, tileType TileType, isSupplyCenter bool) {
	this.Provinces[key] = &Province{Key: key, Name: name, Type: tileType, IsSupplyCenter: isSupplyCenter, Edges: map[string]*Edge{}}
}

func (this *Graph) GetProvince(key string) (*Province, error) {
	if _, ok := this.Provinces[key]; !ok {
		return nil, errors.New("Province not found")
	}

	return this.Provinces[key], nil
}

func (this *Graph) AddEdge(srcKey, destKey string) {
	if _, ok := this.Provinces[srcKey]; !ok {
		return
	}
	if _, ok := this.Provinces[destKey]; !ok {
		return
	}

	this.Provinces[srcKey].Edges[destKey] = &Edge{Province: this.Provinces[destKey]}
}

func (this *Graph) AddUnit(country string, unitType UnitType, province string) error {
	if _, ok := this.Provinces[province]; !ok {
		return errors.New("Province not found")
	}

	if this.Provinces[province].Unit != nil {
		return errors.New("Province already occupied")
	}

	unit := Unit{
		Country: country,
		Type:    unitType,
	}

	this.Provinces[province].Unit = &unit

	return nil
}

func (this *Graph) GetUnits(country string) []*Unit {
	units := []*Unit{}
	for _, tile := range this.Provinces {
		if tile.Unit.Country == country {
			units = append(units, tile.Unit)
		}
	}
	return units
}

func (this *Graph) AddEdges(srcKey string, destKeys []string) {
	if _, ok := this.Provinces[srcKey]; !ok {
		return
	}

	for _, destKey := range destKeys {
		if _, ok := this.Provinces[destKey]; !ok {
			return
		}

		this.Provinces[srcKey].Edges[destKey] = &Edge{Province: this.Provinces[destKey]}
	}
}

func (this *Graph) GetNeighbors(srcKey string) []string {
	result := []string{}

	for _, edge := range this.Provinces[srcKey].Edges {
		result = append(result, edge.Province.Name)
	}

	return result
}
