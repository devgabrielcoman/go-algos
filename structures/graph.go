package structures

import (
	"golang.org/x/exp/slices"
)

type vertex struct {
	value    string
	adjacent []*vertex
}

func Vertex(value string) *vertex {
	return &vertex{
		value:    value,
		adjacent: []*vertex{},
	}
}

func (v *vertex) GetValue() string {
	return v.value
}

func (v *vertex) GetAdjecent() []*vertex {
	return v.adjacent
}

func (v *vertex) Add(adjacent *vertex) {
	if slices.Contains(v.adjacent, adjacent) {
		return
	}
	v.adjacent = append(v.adjacent, adjacent)
	adjacent.Add(v)
}

type graph struct{}

func Graph() *graph {
	return &graph{}
}

func (g *graph) AddVertex(value string) *vertex {
	return Vertex(value)
}

func (g *graph) AddEdge(start *vertex, end *vertex) {
	start.Add(end)
}

func (g *graph) ToArray(start *vertex) []*vertex {
	visited := make(map[*vertex]bool)
	g.depth_first_search(start, visited)

	var result = []*vertex{}
	for key, _ := range visited {
		result = append(result, key)
	}
	return result
}

func (g *graph) depth_first_search(current *vertex, visited map[*vertex]bool) {
	if current == nil {
		return
	}

	// end condition
	_, exists := visited[current]
	if exists {
		return
	}

	// mark the current node as visited
	visited[current] = true

	// iterate over all the adjacent nodes
	for _, adjacent := range current.adjacent {
		g.depth_first_search(adjacent, visited)
	}
}
