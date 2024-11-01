package json

import (
	"encoding/json"
	"os"
	"webbfs/internal/graph"
)

func Export(G *graph.Graph, path string) error {
	data, err := json.MarshalIndent(G, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
