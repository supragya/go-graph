package gograph_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	gograph "github.com/supragya/gograph"
)

func TestCreation(t *testing.T) {
	graph := gograph.NewGraphStringUintString(false)
	assert.Equal(t, graph.GetVertexCount(), 0, "vertex count")
	assert.Equal(t, graph.GetEdgeCount(), 0, "edge count")
}

func TestAddOneEdge(t *testing.T) {
	graph := gograph.NewGraphStringUintString(false)
	graph.AddEdge("vertex1", "vertex2", 1, "metadata")
	assert.Equal(t, graph.GetVertexCount(), 2, "vertex count")
	assert.Equal(t, graph.GetEdgeCount(), 1, "edge count")
	connections := graph.GetConnectedVertices("vertex1")
	assert.Equal(t, len(connections), 1, "number of connections")
	for vertex, edge := range connections {
		assert.Equal(t, vertex, "vertex2", "vertex")
		assert.Equal(t, edge, gograph.WeightedEdge[uint64, string]{1, "metadata"}, "edge")
	}
	connections = graph.GetConnectedVertices("vertex2")
	assert.Equal(t, len(connections), 0, "number of connections")
}

func TestAddOneEdgeBidirectional(t *testing.T) {
	graph := gograph.NewGraphStringUintString(true)
	graph.AddEdge("vertex1", "vertex2", 1, "metadata")
	assert.Equal(t, graph.GetVertexCount(), 2, "vertex count")
	assert.Equal(t, graph.GetEdgeCount(), 1, "edge count")
	connections := graph.GetConnectedVertices("vertex1")
	assert.Equal(t, len(connections), 1, "number of connections")
	for vertex, edge := range connections {
		assert.Equal(t, vertex, "vertex2", "vertex")
		assert.Equal(t, edge, gograph.WeightedEdge[uint64, string]{1, "metadata"}, "edge")
	}
	connections = graph.GetConnectedVertices("vertex2")
	assert.Equal(t, len(connections), 1, "number of connections")
	for vertex, edge := range connections {
		assert.Equal(t, vertex, "vertex1", "vertex")
		assert.Equal(t, edge, gograph.WeightedEdge[uint64, string]{1, "metadata"}, "edge")
	}
}

func TestAPSP1(t *testing.T) {
	graph := gograph.NewGraphStringUintString(false)
	graph.AddEdge("v0", "v1", 1, "v0v1")
	graph.AddEdge("v1", "v2", 1, "v1v2")
	graph.AddEdge("v0", "v2", 3, "v0v2")
	graph.AddEdge("v1", "v3", 1, "v1v3")
	graph.AddEdge("v2", "v3", 1, "v2v3")

	r1 := graph.GetShortestRoute("v0", "v2")
	assert.Equal(t, r1, gograph.Route[string, string]{
		[]string{"v0", "v1", "v2"},
		[]gograph.WeightedEdge[uint64, string]{{1, "v0v1"}, {1, "v1v2"}},
		2,
	}, "shortest route")

	r2 := graph.GetShortestRoute("v1", "v3")
	assert.Equal(t, r2, gograph.Route[string, string]{
		[]string{"v1", "v3"},
		[]gograph.WeightedEdge[uint64, string]{{1, "v1v3"}},
		1,
	}, "shortest route")
}

func TestSaveToDisk(t *testing.T) {
	graph := gograph.NewGraphStringUintString(false)
	graph.AddEdge("v0", "v1", 1, "v0v1")
	graph.AddEdge("v1", "v2", 1, "v1v2")
	graph.AddEdge("v0", "v2", 3, "v0v2")
	graph.AddEdge("v1", "v3", 1, "v1v3")
	graph.AddEdge("v2", "v3", 1, "v2v3")

	graph.CalculateAllPairShortestPath()
	fmt.Print(reflect.TypeOf(graph.AllPairShortestPaths))

	fileName := "/tmp/gograph-test-" + fmt.Sprint(rand.Uint64()) + ".json"
	err := graph.SaveToDisk(fileName)
	assert.Nil(t, err, "save to disk")
}

func TestJSONMarshalling(t *testing.T) {
	route := map[string]map[string]gograph.Route[string, string]{}
	_, err := json.Marshal(route)
	assert.Nil(t, err, "JSON marshalling")
}
