// package graph ...
package graph

import (
	"sync"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

// node type
const (
	PNode = iota
	LNode
	CNode
	ANode

	// DefaultCost default cost.
	DefaultCost = 1
)

// Graph represents a graph with a geometry
// of vertices and weighted edges that can be added or removed.
// The implementation uses hash maps to associate each vertex in the graph with
// its adjacent vertices. This gives constant time performance for
// all basic operations.
type Graph interface {

	// Nodes Returns nodes.
	Nodes() []*Node

	// Edges Returns edges.
	Edges() []map[int]int

	// AddNode add a node.
	AddNode(n *Node)

	// AddNodeType add a node with type.
	AddNodeType(n *Node, nodeType int)

	// AddEdge add a edge.
	AddEdge(n1, n2 *Node)

	// AddEdgeCost add a edge.
	AddEdgeCost(n1, n2 *Node, value int)

	// DeleteNode removes a node .
	DeleteNode(n *Node)

	// DeleteEdge removes an edge from n1 to n2.
	DeleteEdge(n1, n2 *Node)

	// Node tells if there is an node n.
	Node(n *Node) (int, bool)

	// Edge tells if there is an edge from n1 to n2.
	Edge(n1, n2 *Node) bool

	// Order returns the number of vertices in the graph.
	Order() int

	// Union  Computes the Union of two Graph.
	Union(graph Graph) (Graph, error)

	// Intersection  Computes the Intersection of two Graph.
	Intersection(graph Graph) (Graph, error)

	// Difference returns a Graph that represents that part of Graph A that does not intersect with Graph B.
	// One can think of this as GraphA - Intersection(A,B).
	Difference(graph Graph) (Graph, error)

	// SymDifference returns a Graph that represents the portions of A and B that do not intersect.
	// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
	//
	// One can think of this as Union(A,B) - Intersection(A,B).
	SymDifference(graph Graph) (Graph, error)
}

// Node represents a node of a graph
type Node struct {
	index    int
	value    matrix.Steric
	stat     bool
	inGraph  bool
	nodeType int
}

// MatrixGraph represents a graph with a geometry
// of vertices and weighted edges that can be added or removed.
// The implementation uses hash maps to associate each vertex in the graph with
// its adjacent vertices. This gives constant time performance for
// all basic operations.
type MatrixGraph struct {
	nodes []*Node
	edges []map[int]int
	lock  sync.RWMutex
}

// AddNode add a node.
func (g *MatrixGraph) AddNode(n *Node) {
	g.AddNodeType(n, n.nodeType)
}

// AddNodeType add a node with type.
func (g *MatrixGraph) AddNodeType(n *Node, nodeType int) {
	g.lock.Lock()
	defer g.lock.Unlock()
	for _, node := range g.nodes {
		if n.value.Equals(node.value) {
			return
		}
	}
	var node *Node
	if n.inGraph {
		node = &Node{0, n.value, true, true, nodeType}
	} else {
		node = n
	}
	node.index = len(g.nodes)
	node.stat = true
	node.inGraph = true
	g.nodes = append(g.nodes, node)
	edges := make(map[int]int)
	g.edges = append(g.edges, edges)
}

// DeleteNode removes a node.
func (g *MatrixGraph) DeleteNode(n *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.nodes[n.index].stat = false
}

// Node tells if there is an node n.
func (g *MatrixGraph) Node(n *Node) (int, bool) {
	for _, node := range g.nodes {
		if n.value.Equals(node.value) && node.stat {
			return node.index, true
		}
	}
	return -1, false
}

// AddEdge add a edge.
func (g *MatrixGraph) AddEdge(n1, n2 *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()
	index1, ok1 := g.Node(n1)
	index2, ok2 := g.Node(n2)
	if ok1 && ok2 {
		g.edges[index1][index2] = DefaultCost
		g.edges[index2][index1] = DefaultCost
	}
}

// AddEdgeCost add a edge.
func (g *MatrixGraph) AddEdgeCost(n1, n2 *Node, value int) {
	g.lock.Lock()
	defer g.lock.Unlock()
	index1, ok1 := g.Node(n1)
	index2, ok2 := g.Node(n2)
	if ok1 && ok2 {
		g.edges[index1][index2] = value
		g.edges[index2][index1] = value
	}
}

// DeleteEdge removes an edge from n1 to n2.
func (g *MatrixGraph) DeleteEdge(n1, n2 *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.edges == nil {
		return
	}

	delete(g.edges[n1.index], n2.index)
	delete(g.edges[n2.index], n1.index)
}

// Edge tells if there is an edge from n1 to n2.
func (g *MatrixGraph) Edge(n1, n2 *Node) bool {
	index1, ok1 := g.Node(n1)
	index2, ok2 := g.Node(n2)
	if ok1 && ok2 {
		if g.edges[index1][index2] >= DefaultCost &&
			g.edges[index2][index1] >= DefaultCost {
			return true
		}
	}
	return false
}

// Order returns the number of vertices in the graph.
func (g *MatrixGraph) Order() int {
	return len(g.nodes)
}

// Nodes Returns nodes.
func (g *MatrixGraph) Nodes() []*Node {
	return g.nodes
}

// Edges Returns edges.
func (g *MatrixGraph) Edges() []map[int]int {
	return g.edges
}

// Union  Computes the Union of two Graph.
func (g *MatrixGraph) Union(graph Graph) (Graph, error) {
	gUnion := &MatrixGraph{}
	for _, node := range g.Nodes() {
		if node.stat {
			gUnion.AddNode(node)
		}
	}
	for _, node := range graph.Nodes() {
		if node.stat {
			if _, ok := g.Node(node); !ok {
				gUnion.AddNode(node)
			}
		}
	}
	for _, node := range gUnion.nodes {
		if !node.stat {
			break
		}
		if i1, ok1 := g.Node(node); ok1 {
			for k, v := range g.edges[i1] {
				gNode := g.nodes[k]
				gUnion.AddEdgeCost(node, gNode, v)
			}
		}

		if i2, ok2 := graph.Node(node); ok2 {
			for k, v := range graph.Edges()[i2] {
				gNode := graph.Nodes()[k]
				gUnion.AddEdgeCost(node, gNode, v)
			}
		}
	}

	return gUnion, nil

}

// Difference returns a Graph that represents that part of Graph A that does not intersect with Graph B.
// One can think of this as GraphA - Intersection(A,B).
func (g *MatrixGraph) Difference(graph Graph) (Graph, error) {
	gDiff := &MatrixGraph{}
	for _, node := range g.Nodes() {
		if !node.stat {
			break
		}
		gDiff.AddNode(node)
	}

	for _, node := range gDiff.nodes {
		i1, _ := g.Node(node)
		if i2, ok := graph.Node(node); ok {
			numEdge := 0
			for k, v := range g.edges[i1] {
				gNode := g.nodes[k]
				if index2, ok := graph.Node(gNode); ok {
					if graph.Edges()[i2][index2] < DefaultCost {
						gDiff.AddEdgeCost(node, gNode, v)
						numEdge++
					}
				}
			}
			if numEdge == 0 {
				node.stat = false
			}
		} else {
			for k, v := range g.edges[i1] {
				gNode := graph.Nodes()[k]
				gDiff.AddEdgeCost(node, gNode, v)
			}
		}
	}

	return gDiff, nil

}

// SymDifference returns a Graph that represents the portions of A and B that do not intersect.
// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
//
// One can think of this as Union(A,B) - Intersection(A,B).
func (g *MatrixGraph) SymDifference(graph Graph) (Graph, error) {
	var err error
	if gUnion, err := g.Union(graph); err == nil {
		if gIntersect, err := g.Intersection(graph); err == nil {
			return gUnion.Difference(gIntersect)
		}
	}

	return nil, err
}

// Intersection  Computes the Intersection of two Graph.
func (g *MatrixGraph) Intersection(graph Graph) (Graph, error) {
	gIntersect := &MatrixGraph{}
	for _, node := range g.Nodes() {
		if _, ok := graph.Node(node); ok {
			gIntersect.AddNode(node)
		}
	}
	for _, node := range graph.Nodes() {
		if _, ok := graph.Node(node); ok {
			gIntersect.AddNode(node)
		}
	}

	for _, node := range gIntersect.nodes {
		i1, _ := g.Node(node)
		i2, _ := graph.Node(node)
		for k, v := range g.edges[i1] {
			gNode := g.nodes[k]
			if index2, ok := graph.Node(gNode); ok {
				if graph.Edges()[i2][index2] >= DefaultCost {
					gIntersect.AddEdgeCost(node, gNode, v)
				}
			}
		}
	}

	return gIntersect, nil
}
