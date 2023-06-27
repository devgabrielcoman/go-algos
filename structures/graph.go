package structures

import (
	"golang.org/x/exp/maps"
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

func (v *vertex) depth_first_traverse(visited map[*vertex]bool) {
	visited[v] = true

	for _, adjacent := range v.adjacent {
		_, exists := visited[adjacent]
		if !exists {
			adjacent.depth_first_traverse(visited)
		}
	}
}

func (v *vertex) depth_first_search(value string, visited map[*vertex]bool) *vertex {
	// found our target
	if v.value == value {
		return v
	}

	visited[v] = true

	for _, adjacent := range v.adjacent {
		_, exists := visited[adjacent]
		if !exists {
			return adjacent.depth_first_search(value, visited)
		}
	}

	return nil
}

type graph struct {
	last *vertex
}

func Graph() *graph {
	return &graph{
		last: nil,
	}
}

func (g *graph) AddVertex(value string) *vertex {
	new := Vertex(value)
	g.last = new
	return new
}

func (g *graph) AddEdge(start *vertex, end *vertex) {
	start.Add(end)
}

func (g *graph) ToArray() []*vertex {
	if g.last == nil {
		return []*vertex{}
	}

	visited := make(map[*vertex]bool)
	g.last.depth_first_traverse(visited)
	return maps.Keys(visited)
}

func (g *graph) Search(value string) *vertex {
	if g.last == nil {
		return nil
	}

	visited := make(map[*vertex]bool)
	return g.last.depth_first_search(value, visited)
}
