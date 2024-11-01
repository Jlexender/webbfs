package main

import (
	"fmt"
	"webbfs/internal/config"
	"webbfs/internal/crawler"
	"webbfs/internal/crawler/json"
	"webbfs/internal/graph"
)

func main() {
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		panic(err)
	}

	seeds, err := json.LoadSeeds("data/seeds.json")
	if err != nil {
		panic(err)
	}

	webCrawler := crawler.NewCrawler(seeds, cfg.Timeout)
	err = webCrawler.Start(cfg.SearchRegexp)
	if err != nil {
		fmt.Println(err)
	}

	printGraph(webCrawler.Graph)
	err = json.Export(webCrawler.Graph, "data/output.json")
	if err != nil {
		panic(err)
	}
}

func printGraph(G *graph.Graph) {
	fmt.Printf("Total nodes: %d\n", len(G.Nodes))
	for _, node := range G.Nodes {
		fmt.Printf("%s -> ", node.Url)
		for _, child := range node.Children {
			fmt.Printf("%s, ", child.Url)
		}
		fmt.Println()
	}
}
