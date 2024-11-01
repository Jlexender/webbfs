package net

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

var client = &http.Client{}
var timeout time.Duration

func Init(timeoutSeconds float32) {
	timeout = time.Duration(timeoutSeconds)
}

func GetResponseBody(url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to make request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
