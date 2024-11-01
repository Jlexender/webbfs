package crawler

import (
	"regexp"
	"webbfs/internal/graph"
	"webbfs/internal/net"
	"webbfs/internal/util"
)

func WebSearch(node *graph.WebNode, G *graph.Graph, queue *util.PriorityQueue[*graph.WebNode], re *regexp.Regexp) {
	body, _ := net.GetResponseBody(node.Url)

	links := re.FindAllString(string(body), -1)
	for _, link := range links {
		if G.Nodes[link] == nil {
			G.Nodes[link] = graph.NewNode(link, node.Depth-1)
			if G.Nodes[link].Depth > 0 {
				queue.Push(G.Nodes[link])
			}
		}
		if G.Nodes[link].Depth < node.Depth-1 {
			G.Nodes[link].Depth = node.Depth - 1
		}
		if !G.Nodes[node.Url].Contains(link) {
			G.Nodes[node.Url].Children = append(G.Nodes[node.Url].Children, G.Nodes[link])
		}
		G.Nodes[node.Url].Edges[link]++
	}
}
