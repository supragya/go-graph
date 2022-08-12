package gograph

type Graph[K comparable, W comparable, V any] struct {
	isBidirectional bool
	vertices        map[K]Connections[K, W, V]
	vertexCount     int
	edgeCount       int
}

func NewGraphStringUintString(bidirectional bool) *Graph[string, uint, string] {
	return &Graph[string, uint, string]{
		isBidirectional: bidirectional,
		vertices:        make(map[string]Connections[string, uint, string]),
		vertexCount:     0,
		edgeCount:       0,
	}
}

func (g *Graph[K, W, V]) ensureVertexAvailable(vertex K) {
	_, isAvailable := g.vertices[vertex]
	if !isAvailable {
		g.vertices[vertex] = make(Connections[K, W, V])
		g.vertexCount++
	}
}

func (g *Graph[K, W, V]) GetConnectedVertices(vertex K) Connections[K, W, V] {
	g.ensureVertexAvailable(vertex)
	connectedVertices, _ := g.vertices[vertex]
	return connectedVertices
}

func (g *Graph[K, W, V]) AddEdge(from K, to K, weight W, edge V) error {
	cFrom := g.GetConnectedVertices(from)
	cTo := g.GetConnectedVertices(to)

	if cFrom.Exists(to) || cTo.Exists(from) {
		return ErrEdgeExists
	}

	g.vertices[from] = *cFrom.Added(to, weight, edge)
	if g.isBidirectional {
		g.vertices[to] = *cTo.Added(from, weight, edge)
	}

	g.edgeCount++

	return nil
}

func (g *Graph[K, W, V]) GetVertexCount() int {
	return g.vertexCount
}

func (g *Graph[K, W, V]) GetEdgeCount() int {
	return g.edgeCount
}
