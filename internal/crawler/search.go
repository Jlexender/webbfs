package crawler

import (
	"log"
	"regexp"
	"urlCrawl/internal/graph"
	"urlCrawl/internal/util"
)

// Search function based on the BFS algorithm. /**
func Search(start *graph.Node, G *graph.Graph, re *regexp.Regexp,
	action func(*graph.Node, *graph.Graph, *util.PriorityQueue[*graph.Node], *regexp.Regexp) error) error {
	queue := util.NewPriorityQueue[*graph.Node]()
	visited := make(map[string]int)

	queue.Push(start)
	visited[start.Url] = start.Depth

	for queue.Len() > 0 {

		node := *queue.Pop()
		if node.Depth > 0 {
			err := action(node, G, queue, re)
			if err != nil {
				return err
			}
		}
	}
	log.Println("Search completed")
	return nil
}
