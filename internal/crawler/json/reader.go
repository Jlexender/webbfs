package json

import (
	"encoding/json"
	"os"
	"webbfs/internal/graph"
)

func LoadSeeds(path string) ([]*graph.WebNode, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var seeds []*graph.WebNode
	err = json.Unmarshal(data, &seeds)
	if err != nil {
		return nil, err
	}
	initSeeds(seeds)

	return seeds, nil
}

func initSeeds(seeds []*graph.WebNode) {
	for _, seed := range seeds {
		seed.Children = make([]*graph.WebNode, 0)
		seed.Edges = make(map[string]int)
	}
}
