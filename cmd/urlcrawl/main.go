package main

import (
	"fmt"
	"regexp"
	"urlCrawl/internal/crawler"
	"urlCrawl/internal/graph"
)

func main() {
	depth := 10

	G := graph.NewGraph()
	start := graph.NewNode("https://www.instagram.com", depth)
	regex := regexp.MustCompile(`https?://([a-zA-Z0-9-]+\.[a-zA-Z0-9.-]+)`)

	G.AddNode(start)
	crawler.Search(start, G, regex, crawler.WebSearch)

	for _, node := range G.Nodes {
		fmt.Printf("%s -> ", node.Url)
		for _, child := range node.Children {
			fmt.Printf("%s, ", child.Url)
		}
		fmt.Println()
	}
}
