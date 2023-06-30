package structures_test

import (
	"gabriel/algos/structures"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func Test_Vertex(t *testing.T) {
	t.Run("can create a new vertex", func(t *testing.T) {
		vertex := structures.NewVertex("john")
		assert.NotNil(t, vertex)
		assert.Equal(t, "john", vertex.Value)
		assert.Equal(t, 0, len(vertex.Adjacent))
	})

	t.Run("a vertex can add another vertex", func(t *testing.T) {
		vertex1 := structures.NewVertex("john")
		vertex2 := structures.NewVertex("mary")

		vertex1.Add(vertex2)

		assert.True(t, slices.Contains(vertex1.Adjacent, vertex2))
		assert.True(t, slices.Contains(vertex2.Adjacent, vertex1))
	})
}

func Test_Graph(t *testing.T) {
	t.Run("can create a new graph", func(t *testing.T) {
		graph := structures.NewGraph()
		assert.NotNil(t, graph)
	})
}

func Test_Graph_AddVertexAndEdges(t *testing.T) {
	t.Run("can create a new graph", func(t *testing.T) {
		graph := structures.NewGraph()
		assert.NotNil(t, graph)
	})

	t.Run("can add vertices", func(t *testing.T) {
		graph := structures.NewGraph()

		vertex1 := graph.AddVertex("john")
		vertex2 := graph.AddVertex("mary")

		assert.NotNil(t, vertex1)
		assert.NotNil(t, vertex2)
	})

	t.Run("can add edges", func(t *testing.T) {
		graph := structures.NewGraph()

		vertex1 := graph.AddVertex("john")
		vertex2 := graph.AddVertex("mary")
		graph.AddEdge(vertex1, vertex2)

		assert.True(t, slices.Contains(vertex1.Adjacent, vertex2))
		assert.True(t, slices.Contains(vertex2.Adjacent, vertex1))
	})
}

func Test_Graph_Traverse(t *testing.T) {
	t.Run("can traverse an empty graph depth first", func(t *testing.T) {
		graph := structures.NewGraph()

		result := pluckValues(graph.TraverseDeep())
		assert.Equal(t, []string{}, result)
	})

	t.Run("can traverse a simple graph depth first", func(t *testing.T) {
		graph := provideSimpleTestGraph()

		result := pluckValues(graph.TraverseDeep())
		assert.Equal(t, []string{"john", "mary"}, result)
	})

	t.Run("can traverse a complex graph depth first", func(t *testing.T) {
		graph := provideComplexTestGraph()

		result := pluckValues(graph.TraverseDeep())
		assert.Equal(t, []string{"alice", "bob", "candy", "derek", "elaine", "fred", "gina", "helen", "irena"}, result)
	})

	t.Run("can traverse an empty graph breath first", func(t *testing.T) {
		graph := structures.NewGraph()

		result := pluckValues(graph.TraverseWide())
		assert.Equal(t, []string{}, result)
	})

	t.Run("can traverse a simple graph breath first", func(t *testing.T) {
		graph := provideSimpleTestGraph()

		result := pluckValues(graph.TraverseWide())
		assert.Equal(t, []string{"john", "mary"}, result)
	})

	t.Run("can traverse a complex graph breath first", func(t *testing.T) {
		graph := provideComplexTestGraph()
		result := pluckValues(graph.TraverseWide())
		assert.Equal(t, []string{"alice", "bob", "candy", "derek", "elaine", "fred", "gina", "helen", "irena"}, result)
	})
}

func Test_Graph_Search(t *testing.T) {
	t.Run("returns nil trying to search depth first in an empty graph", func(t *testing.T) {
		graph := structures.NewGraph()
		result := graph.SearchDeep("alice")
		assert.Nil(t, result)
	})

	t.Run("returns nil trying to search depth first for a non existing value", func(t *testing.T) {
		graph := provideSimpleTestGraph()

		result := graph.SearchDeep("bob")
		assert.Nil(t, result)
	})

	t.Run("returns the correct vertex trying to search depth first for an existing value", func(t *testing.T) {
		graph := provideComplexTestGraph()

		result := graph.SearchDeep("derek")
		assert.NotNil(t, result)
		assert.Equal(t, "derek", result.Value)
	})

	t.Run("returns nil trying to search breadth first in an empty graph", func(t *testing.T) {
		graph := structures.NewGraph()
		result := graph.SearchWide("alice")
		assert.Nil(t, result)
	})

	t.Run("returns nil trying to search breadth first for a non existing value", func(t *testing.T) {
		graph := provideSimpleTestGraph()

		result := graph.SearchWide("bob")
		assert.Nil(t, result)
	})

	t.Run("returns the correct vertex trying to search breadth first for an existing value", func(t *testing.T) {
		graph := provideComplexTestGraph()

		result := graph.SearchWide("derek")
		assert.NotNil(t, result)
		assert.Equal(t, "derek", result.Value)
	})
}

func provideSimpleTestGraph() *structures.Graph {
	graph := structures.NewGraph()

	vertex1 := graph.AddVertex("john")
	vertex2 := graph.AddVertex("mary")
	graph.AddEdge(vertex1, vertex2)
	return graph
}

func provideComplexTestGraph() *structures.Graph {
	graph := structures.NewGraph()

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

	return graph
}

func pluckValues(verices []*structures.Vertex) []string {
	var result = []string{}

	for _, vertex := range verices {
		result = append(result, vertex.Value)
	}

	slices.Sort(result)

	return result
}
