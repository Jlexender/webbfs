package main

import (
	"fmt"
	"webbfs/internal/config"
	"webbfs/internal/crawler"
	"webbfs/internal/crawler/json"
	"webbfs/internal/graph"
	"webbfs/internal/net"
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

	net.Init(cfg.NetTimeoutSeconds)

	newCrawler := crawler.NewCrawler(seeds)
	err = newCrawler.Start(cfg.SearchRegexp)
	if err != nil {
		fmt.Println(err)
	}

	printGraph(newCrawler.Graph)
	err = json.Export(newCrawler.Graph, "data/output.json")
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
