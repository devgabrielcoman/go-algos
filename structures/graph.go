package structures

import (
	"golang.org/x/exp/maps"
)

type Vertex struct {
	Value    string
	Adjacent []*Vertex
}

func NewVertex(value string) *Vertex {
	return &Vertex{
		Value:    value,
		Adjacent: []*Vertex{},
	}
}

func (v *Vertex) Add(adjacent *Vertex) {
	v.Adjacent = append(v.Adjacent, adjacent)
	adjacent.Adjacent = append(adjacent.Adjacent, v)
}

func (v *Vertex) traverse_depth_first(visited map[*Vertex]bool) []*Vertex {
	visited[v] = true

	for _, adjacent := range v.Adjacent {
		_, exists := visited[adjacent]
		if !exists {
			adjacent.traverse_depth_first(visited)
		}
	}

	return maps.Keys(visited)
}

func (v *Vertex) search_depth_first(value string, visited map[*Vertex]bool) *Vertex {
	if v.Value == value {
		return v
	}

	visited[v] = true

	for _, adjacent := range v.Adjacent {
		_, exists := visited[adjacent]
		if !exists {
			return adjacent.search_depth_first(value, visited)
		}
	}

	return nil
}

func (v *Vertex) traverse_breadth_first(visited map[*Vertex]bool) []*Vertex {

	visited[v] = true

	queue := Queue[Vertex]()
	queue.Enqueue(*v)

	for !queue.IsEmpty() {
		current := queue.Dequeue()
		for _, adjacent := range current.Adjacent {
			_, exists := visited[adjacent]
			if !exists {
				visited[adjacent] = true
				queue.Enqueue(*adjacent)
			}
		}
	}

	return maps.Keys(visited)
}

func (v *Vertex) searcH_breadth_first(value string, visited map[*Vertex]bool) *Vertex {
	visited[v] = true

	queue := Queue[Vertex]()
	queue.Enqueue(*v)

	for !queue.IsEmpty() {
		current := queue.Dequeue()

		if current.Value == value {
			return current
		}

		for _, adjacent := range current.Adjacent {
			_, exists := visited[adjacent]
			if !exists {
				visited[adjacent] = true
				queue.Enqueue(*adjacent)
			}
		}
	}

	return nil
}

type Graph struct {
	last *Vertex
}

func NewGraph() *Graph {
	return &Graph{
		last: nil,
	}
}

func (g *Graph) AddVertex(value string) *Vertex {
	new := NewVertex(value)
	g.last = new
	return new
}

func (g *Graph) AddEdge(start *Vertex, end *Vertex) {
	start.Add(end)
}

func (g *Graph) TraverseDeep() []*Vertex {
	if g.last == nil {
		return []*Vertex{}
	}

	visited := make(map[*Vertex]bool)
	return g.last.traverse_depth_first(visited)
}

func (g *Graph) SearchDeep(value string) *Vertex {
	if g.last == nil {
		return nil
	}

	visited := make(map[*Vertex]bool)
	return g.last.search_depth_first(value, visited)
}

func (g *Graph) TraverseWide() []*Vertex {
	if g.last == nil {
		return []*Vertex{}
	}

	visited := make(map[*Vertex]bool)
	return g.last.traverse_breadth_first(visited)
}

func (g *Graph) SearchWide(value string) *Vertex {
	if g.last == nil {
		return nil
	}

	visited := make(map[*Vertex]bool)
	return g.last.searcH_breadth_first(value, visited)
}
