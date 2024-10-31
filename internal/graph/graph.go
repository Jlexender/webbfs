package graph

type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}

func (g *Graph) AddNode(node *Node) {
	g.Nodes[node.Url] = node
}
