package crawler

import (
	"regexp"
	"webbfs/internal/graph"
	"webbfs/internal/util"
)

// Search function based on the BFS algorithm. /**
func Search(start *graph.WebNode, G *graph.Graph, re *regexp.Regexp,
	action func(*graph.WebNode, *graph.Graph, *util.PriorityQueue[*graph.WebNode], *regexp.Regexp)) {
	queue := util.NewPriorityQueue[*graph.WebNode]()
	visited := make(map[string]int)

	queue.Push(start)
	visited[start.Url] = start.Depth

	for queue.Len() > 0 {

		node := *queue.Pop()
		if node.Depth > 0 {
			action(node, G, queue, re)
		}
	}
}
