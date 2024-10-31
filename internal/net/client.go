package net

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

var client = &http.Client{}
var timeout time.Duration = 3

func GetResponseBody(url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error fetching response body for URL: %s, status code: %d", url, resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
