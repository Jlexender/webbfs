package crawler

import (
	"regexp"
	"webbfs/internal/graph"
)

type Crawler struct {
	Seeds []*graph.WebNode
	Graph *graph.Graph
}

func NewCrawler(seeds []*graph.WebNode) *Crawler {
	return &Crawler{
		Seeds: seeds,
		Graph: graph.NewGraph(),
	}
}

func (c *Crawler) Start(re string) error {
	regex := regexp.MustCompile(re)
	for _, seed := range c.Seeds {
		c.Graph.AddNode(seed)
		Search(seed, c.Graph, regex, WebSearch)
	}
	return nil
}
