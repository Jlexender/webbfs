package crawler

import (
	"regexp"
	"webbfs/internal/graph"
	"webbfs/internal/net"
	"webbfs/internal/util"
)

type SearchRequest struct {
	Start  *graph.WebNode
	Graph  *graph.Graph
	Re     *regexp.Regexp
	Action func(*graph.WebNode, *graph.Graph, *util.PriorityQueue[*graph.WebNode], *regexp.Regexp, *net.Client)
	Client *net.Client
}

// Search function based on the BFS algorithm. /**
func Search(req SearchRequest) {
	queue := util.NewPriorityQueue[*graph.WebNode]()
	visited := make(map[string]int)

	queue.Push(req.Start)
	visited[req.Start.Url] = req.Start.Depth

	for queue.Len() > 0 {

		node := *queue.Pop()
		if node.Depth > 0 {
			req.Action(node, req.Graph, queue, req.Re, req.Client)
		}
	}
}
