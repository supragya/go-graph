package gograph

type Graph[K comparable, V comparable] struct {
	isBidirectional bool
	vertices        map[K]Connections[K, V]
	vertexCount     int
	edgeCount       int
}

func NewGraphStringVertexStringEdge(bidirectional bool) *Graph[string, string] {
	return &Graph[string, string]{
		isBidirectional: bidirectional,
		vertices:        make(map[string]Connections[string, string]),
		vertexCount:     0,
		edgeCount:       0,
	}
}

func (g *Graph[K, V]) ensureVertexAvailable(vertex K) {
	_, isAvailable := g.vertices[vertex]
	if !isAvailable {
		g.vertices[vertex] = make(Connections[K, V])
		g.vertexCount++
	}
}

func (g *Graph[K, V]) GetConnectedVertices(vertex K) Connections[K, V] {
	g.ensureVertexAvailable(vertex)
	connectedVertices, _ := g.vertices[vertex]
	return connectedVertices
}

func (g *Graph[K, V]) AddEdge(from K, to K, edge V) error {
	cFrom := g.GetConnectedVertices(from)
	cTo := g.GetConnectedVertices(to)

	if cFrom.Exists(to) || cTo.Exists(from) {
		return ErrEdgeExists
	}

	g.vertices[from] = *cFrom.Added(to, edge)
	if g.isBidirectional {
		g.vertices[to] = *cTo.Added(from, edge)
	}

	g.edgeCount++

	return nil
}

func (g *Graph[K, V]) GetVertexCount() int {
	return g.vertexCount
}

func (g *Graph[K, V]) GetEdgeCount() int {
	return g.edgeCount
}
