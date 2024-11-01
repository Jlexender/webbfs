package graph

type Graph struct {
	Nodes map[string]*WebNode
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*WebNode),
	}
}

func (g *Graph) AddNode(node *WebNode) {
	g.Nodes[node.Url] = node
}
