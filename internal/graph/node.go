package graph

type WebNode struct {
	Url      string         `json:"url"`
	Depth    int            `json:"depth"`
	Children []*WebNode     `json:"-"`
	Edges    map[string]int `json:"edges"`
}

func (n *WebNode) Compare(other *WebNode) bool {
	return n.Depth > other.Depth
}

func NewNode(url string, depth int) *WebNode {
	return &WebNode{
		Url:      url,
		Depth:    depth,
		Children: []*WebNode{},
		Edges:    make(map[string]int),
	}
}

func (n *WebNode) Contains(url string) bool {
	return n.Edges[url] != 0
}
