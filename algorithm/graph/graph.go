// package graph ...

package graph

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

// node type
const (
	PNode = 1
	LNode = 2
	CNode = 4
	ANode = 8

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

	// Degree Returns degrees(num Connected) of node.
	Degree(index int) int

	// Connected Returns degree of node.
	Connected(index int) int

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

	// Node tells if there is an node .
	Node(n *Node) (*Node, bool)

	// NodeIndex tells if there is an node index.
	NodeIndex(n *Node) (int, bool)

	// Edge tells if there is an edge from n1 to n2.
	Edge(n1, n2 *Node) bool

	// Order returns the number of vertices in the graph.
	Order() int

	// Equals returns the true if g==g1.
	Equals(g1 Graph) bool

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

	// String ...
	String() string
}

// Node represents a node of a graph
type Node struct {
	Index    int
	Value    matrix.Steric
	Reverse  matrix.Steric
	stat     bool
	inGraph  bool
	NodeType int
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

// String ...
func (g *MatrixGraph) String() string {
	str := ""
	for i, v := range g.nodes {
		str += fmt.Sprintf("node %v: %v", i, v.Value)
	}
	return fmt.Sprintf("Nodes:%v\nEdges%v", str, g.edges)
}

// Equals returns the true if g==g1.
func (g *MatrixGraph) Equals(g1 Graph) bool {
	if g1.Order() != g.Order() {
		return false
	}
	for i := 0; i < g1.Order(); i++ {
		if g.Nodes()[i].stat {
			if !g.Nodes()[i].Value.Equals(g1.Nodes()[i].Value) &&
				!g.Nodes()[i].Reverse.Equals(g1.Nodes()[i].Value) {
				return false
			}
		}
	}
	for i := 0; i < g1.Order(); i++ {
		if g.Nodes()[i].stat {
			if !reflect.DeepEqual(g.Edges()[i], g1.Edges()[i]) {
				return false
			}
		}
	}

	return true
}

// AddNode add a node.
func (g *MatrixGraph) AddNode(n *Node) {
	g.AddNodeType(n, n.NodeType)
}

// AddNodeType add a node with type.
func (g *MatrixGraph) AddNodeType(n *Node, nodeType int) {
	g.lock.Lock()
	defer g.lock.Unlock()
	for _, node := range g.nodes {
		if n.Value.Equals(node.Value) || n.Value.Equals(node.Reverse) {
			return
		}
	}
	var node *Node
	if n.inGraph {
		node = &Node{Index: 0, Value: n.Value, Reverse: n.Reverse, stat: true, inGraph: true, NodeType: nodeType}
	} else {
		node = n
	}
	if v, ok := node.Value.(matrix.LineMatrix); ok {
		rl := matrix.LineMatrix{}
		for i := len(v) - 1; i >= 0; i-- {
			rl = append(rl, v[i])
		}
		n.Reverse = rl
	} else {
		n.Reverse = n.Value
	}
	node.Index = len(g.nodes)
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

	g.nodes[n.Index].stat = false
	for k := range g.edges[n.Index] {
		g.DeleteEdgeByIndex(k, n.Index)
	}
}

// Node tells if there is an node .
func (g *MatrixGraph) Node(n *Node) (*Node, bool) {
	for _, node := range g.nodes {
		if n.Value.Equals(node.Value) && node.stat ||
			n.Value.Equals(node.Reverse) && node.stat {
			return node, true
		}
	}
	return nil, false
}

// NodeIndex tells if there is an node index.
func (g *MatrixGraph) NodeIndex(n *Node) (int, bool) {
	if node, ok := g.Node(n); ok {
		return node.Index, true
	}
	return -1, false
}

// AddEdge add a edge.
func (g *MatrixGraph) AddEdge(n1, n2 *Node) {
	g.AddEdgeCost(n1, n2, calcCost(n1, n2))
}

// AddEdgeCost add a edge.
func (g *MatrixGraph) AddEdgeCost(n1, n2 *Node, value int) {
	g.lock.Lock()
	defer g.lock.Unlock()
	node1, ok1 := g.Node(n1)
	node2, ok2 := g.Node(n2)
	if ok1 && ok2 {
		g.edges[node1.Index][node2.Index] = value
		g.edges[node2.Index][node1.Index] = value
	}
}

func calcCost(n1, n2 *Node) int {
	return n1.NodeType + n2.NodeType
}

// DeleteEdge removes an edge from n1 to n2.
func (g *MatrixGraph) DeleteEdge(n1, n2 *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.DeleteEdgeByIndex(n1.Index, n2.Index)
}

// DeleteEdgeByIndex removes an edge from n1 to n2.
func (g *MatrixGraph) DeleteEdgeByIndex(n1, n2 int) {
	if g.edges == nil {
		return
	}

	delete(g.edges[n1], n2)
	delete(g.edges[n2], n1)
}

// Edge tells if there is an edge from n1 to n2.
func (g *MatrixGraph) Edge(n1, n2 *Node) bool {
	index1, ok1 := g.NodeIndex(n1)
	index2, ok2 := g.NodeIndex(n2)
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
	num := 0
	for _, v := range g.nodes {
		if v.stat {
			num++
		}
	}
	return num
}

// Nodes Returns nodes.
func (g *MatrixGraph) Nodes() []*Node {
	return g.nodes
}

// Edges Returns edges.
func (g *MatrixGraph) Edges() []map[int]int {
	return g.edges
}

// Degree Returns degree of node.
func (g *MatrixGraph) Degree(index int) int {
	return len(g.Edges()[index])
}

// Connected Returns num Connected of node.
func (g *MatrixGraph) Connected(index int) int {
	switch g.Nodes()[index].NodeType {
	case PNode:
		if g.Degree(index) == 0 {
			return 0
		}
		return 1
	default:
		if g.Degree(index) == 2 {
			return 2
		}
		return 1
	}
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
		if i1, ok1 := g.NodeIndex(node); ok1 {
			for k, v := range g.edges[i1] {
				gNode := g.nodes[k]
				gUnion.AddEdgeCost(node, gNode, v)
			}
		}

		if i2, ok2 := graph.NodeIndex(node); ok2 {
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
		i1, _ := g.NodeIndex(node)
		if i2, ok := graph.NodeIndex(node); ok {
			numEdge := 0
			for k, v := range g.edges[i1] {
				gNode := g.nodes[k]
				if index2, ok := graph.NodeIndex(gNode); ok {
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

	for _, node := range gIntersect.nodes {
		if i1, ok1 := g.NodeIndex(node); ok1 {
			i2, _ := graph.NodeIndex(node)
			if g.edges == nil || g.edges[i1] == nil {
				continue
			}
			for k, v := range g.edges[i1] {
				gNode := g.nodes[k]
				if index2, ok := graph.NodeIndex(gNode); ok {
					if graph.Edges()[i2][index2] >= DefaultCost {
						gIntersect.AddEdgeCost(node, gNode, v)
					}
				}
			}
		}
	}

	return gIntersect, nil
}
