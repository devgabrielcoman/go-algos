package structures_test

import (
	"gabriel/algos/structures"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func Test_Vertex(t *testing.T) {
	t.Run("can create a new vertex", func(t *testing.T) {
		vertex := structures.Vertex("john")
		assert.NotNil(t, vertex)
		assert.Equal(t, "john", vertex.GetValue())
		assert.Equal(t, 0, len(vertex.GetAdjecent()))
	})

	t.Run("a vertex can add another vertex", func(t *testing.T) {
		vertex1 := structures.Vertex("john")
		vertex2 := structures.Vertex("mary")

		vertex1.Add(vertex2)

		assert.True(t, slices.Contains(vertex1.GetAdjecent(), vertex2))
		assert.True(t, slices.Contains(vertex2.GetAdjecent(), vertex1))
	})
}

func Test_Graph(t *testing.T) {
	t.Run("can create a new graph", func(t *testing.T) {
		graph := structures.Graph()
		assert.NotNil(t, graph)
	})

	t.Run("can add vertices", func(t *testing.T) {
		graph := structures.Graph()

		vertex1 := graph.AddVertex("john")
		vertex2 := graph.AddVertex("mary")

		assert.NotNil(t, vertex1)
		assert.NotNil(t, vertex2)
	})

	t.Run("can add edges", func(t *testing.T) {
		graph := structures.Graph()

		vertex1 := graph.AddVertex("john")
		vertex2 := graph.AddVertex("mary")

		graph.AddEdge(vertex1, vertex2)

		assert.True(t, slices.Contains(vertex1.GetAdjecent(), vertex2))
		assert.True(t, slices.Contains(vertex2.GetAdjecent(), vertex1))
	})

	t.Run("can convert an empty graph to an empty array", func(t *testing.T) {
		graph := structures.Graph()
		result := graph.ToArray()
		assert.Equal(t, 0, len(result))
	})

	t.Run("can convert a simple graph to an array", func(t *testing.T) {
		graph := structures.Graph()
		vertex1 := graph.AddVertex("john")
		vertex2 := graph.AddVertex("mary")
		graph.AddEdge(vertex1, vertex2)

		result := graph.ToArray()
		assert.Equal(t, 2, len(result))
	})

	t.Run("can convert complex graph to an array", func(t *testing.T) {
		graph := structures.Graph()

		alice := graph.AddVertex("alice")
		bob := graph.AddVertex("bob")
		fred := graph.AddVertex("fred")
		helen := graph.AddVertex("helen")
		candy := graph.AddVertex("candy")
		derek := graph.AddVertex("derek")
		gina := graph.AddVertex("gina")
		elaine := graph.AddVertex("elaine")
		irena := graph.AddVertex("irena")

		graph.AddEdge(alice, bob)
		graph.AddEdge(alice, candy)
		graph.AddEdge(alice, derek)
		graph.AddEdge(alice, elaine)
		graph.AddEdge(bob, fred)
		graph.AddEdge(fred, helen)
		graph.AddEdge(helen, candy)
		graph.AddEdge(derek, elaine)
		graph.AddEdge(derek, gina)
		graph.AddEdge(gina, irena)

		result := graph.ToArray()
		assert.Equal(t, 9, len(result))
	})

	t.Run("returns nil trying to search in an empty graph", func(t *testing.T) {
		graph := structures.Graph()
		result := graph.Search("alice")
		assert.Nil(t, result)
	})

	t.Run("returns nil trying to search for a non existing value", func(t *testing.T) {
		graph := structures.Graph()
		vertex1 := graph.AddVertex("john")
		vertex2 := graph.AddVertex("mary")
		graph.AddEdge(vertex1, vertex2)

		result := graph.Search("bob")
		assert.Nil(t, result)
	})

	t.Run("returns the correct vertex trying to search for an existing value", func(t *testing.T) {
		graph := structures.Graph()
		vertex1 := graph.AddVertex("john")
		vertex2 := graph.AddVertex("mary")
		graph.AddEdge(vertex1, vertex2)

		result := graph.Search("mary")
		assert.Equal(t, result, vertex2)
	})
}
