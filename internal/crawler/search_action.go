package crawler

import (
	"regexp"
	"urlCrawl/internal/graph"
	"urlCrawl/internal/net"
	"urlCrawl/internal/util"
)

func WebSearch(node *graph.Node, G *graph.Graph, queue *util.PriorityQueue[*graph.Node], re *regexp.Regexp) error {
	body, err := net.GetResponseBody(node.Url)
	if err != nil {
		return err
	}

	links := util.FindAllMatches(body, re)
	for _, link := range links {
		if G.Nodes[link] == nil {
			G.Nodes[link] = graph.NewNode(link, node.Depth-1)
			G.Nodes[node.Url].Children = append(G.Nodes[node.Url].Children, G.Nodes[link])
			if G.Nodes[link].Depth > 0 {
				queue.Push(G.Nodes[link])
			}
		}
		if G.Nodes[link].Depth < node.Depth-1 {
			G.Nodes[link].Depth = node.Depth - 1
		}
	}

	return nil
}
