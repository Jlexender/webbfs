package graph

type Node struct {
	Url      string
	Depth    int
	Children []*Node
}

func (n *Node) Compare(other *Node) bool {
	return n.Depth > other.Depth
}

func NewNode(url string, depth int) *Node {
	return &Node{
		Url:      url,
		Depth:    depth,
		Children: []*Node{},
	}
}
