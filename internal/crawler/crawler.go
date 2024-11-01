package crawler

import (
	"regexp"
	"webbfs/internal/graph"
	"webbfs/internal/net"
)

type Crawler struct {
	Seeds   []*graph.WebNode
	Graph   *graph.Graph
	Timeout float32
}

func NewCrawler(seeds []*graph.WebNode, timeout float32) *Crawler {
	return &Crawler{
		Seeds:   seeds,
		Graph:   graph.NewGraph(),
		Timeout: timeout,
	}
}

func (c *Crawler) Start(re string) error {
	regex := regexp.MustCompile(re)

	for _, seed := range c.Seeds {
		c.Graph.AddNode(seed)
		req := SearchRequest{
			Start:  seed,
			Graph:  c.Graph,
			Re:     regex,
			Action: WebSearch,
			Client: net.NewClient(c.Timeout),
		}
		Search(req)
	}
	return nil
}
