package gograph

type Connections[K comparable, V any] map[K]V

func (s *Connections[K, V]) Exists(item K) bool {
	if s == nil {
		return false
	}
	_, exists := (*s)[item]
	return exists
}

func (s *Connections[K, V]) Added(item K, edge V) *Connections[K, V] {
	if s == nil {
		sNew := make(Connections[K, V])
		sNew[item] = edge
		return &sNew
	}
	(*s)[item] = edge
	return s
}
