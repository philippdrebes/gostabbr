package engine

type TileType int8

const (
	LandTile TileType = iota
	WaterTile
)

type Graph struct {
	Vertices map[string]*Vertex
}

type Vertex struct {
	Key   string
	Name  string
	Type  TileType
	Edges map[string]*Edge
}

type Edge struct {
	Vertex *Vertex
}

func (this *Graph) AddVertex(key, name string, tileType TileType) {
	this.Vertices[key] = &Vertex{Key: key, Name: name, Type: tileType, Edges: map[string]*Edge{}}
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

func (this *Graph) GetNeighbors(srcKey string) []string {
	result := []string{}

	for _, edge := range this.Vertices[srcKey].Edges {
		result = append(result, edge.Vertex.Name)
	}

	return result
}
