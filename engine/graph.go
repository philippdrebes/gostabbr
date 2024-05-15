package engine

import (
	"errors"
	"fmt"
)

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
	Order   *Order
	Country *Country
	Type    UnitType
}

func (g *Graph) AddProvince(key, name string, tileType TileType, isSupplyCenter bool) {
	g.Provinces[key] = &Province{Key: key, Name: name, Type: tileType, IsSupplyCenter: isSupplyCenter, Edges: map[string]*Edge{}}
}

func (g *Graph) GetProvince(key string) (*Province, error) {
	if _, ok := g.Provinces[key]; !ok {
		return nil, errors.New(fmt.Sprintf("Province '%s' not found", key))
	}

	return g.Provinces[key], nil
}

func (g *Graph) AddEdge(srcKey, destKey string) {
	if _, ok := g.Provinces[srcKey]; !ok {
		return
	}
	if _, ok := g.Provinces[destKey]; !ok {
		return
	}

	g.Provinces[srcKey].Edges[destKey] = &Edge{Province: g.Provinces[destKey]}
}

func (g *Graph) AddUnit(country *Country, unitType UnitType, province string) (*Unit, error) {
	if _, ok := g.Provinces[province]; !ok {
		return nil, errors.New("Province not found")
	}

	if g.Provinces[province].Unit != nil {
		return nil, errors.New("Province already occupied")
	}

	unit := &Unit{
		Country: country,
		Type:    unitType,
	}

	g.Provinces[province].Unit = unit

	return unit, nil
}

func (g *Graph) GetUnits(country string) []*Unit {
	units := []*Unit{}
	for _, tile := range g.Provinces {
		if tile.Unit.Country.Name == country {
			units = append(units, tile.Unit)
		}
	}
	return units
}

func (g *Graph) AddEdges(srcKey string, destKeys []string) {
	if _, ok := g.Provinces[srcKey]; !ok {
		return
	}

	for _, destKey := range destKeys {
		if _, ok := g.Provinces[destKey]; !ok {
			return
		}

		g.Provinces[srcKey].Edges[destKey] = &Edge{Province: g.Provinces[destKey]}
	}
}

func (g *Graph) GetNeighborKeys(srcKey string) []string {
	result := []string{}

	for _, edge := range g.Provinces[srcKey].Edges {
		result = append(result, edge.Province.Name)
	}

	return result
}

func (g *Graph) GetNeighbors(src *Province) ([]*Province, error) {
	if src == nil {
		return nil, errors.New("input is nil")
	}

	result := []*Province{}

	for _, edge := range g.Provinces[src.Key].Edges {
		result = append(result, edge.Province)
	}

	return result, nil
}

func (g *Graph) GetNeighborsWithUnits(src *Province) ([]*Province, error) {
	if src == nil {
		return nil, errors.New("input is nil")
	}

	result := []*Province{}

	for _, edge := range g.Provinces[src.Key].Edges {
		p := edge.Province
		if p.Unit != nil {
			result = append(result, p)
		}
	}

	return result, nil
}
