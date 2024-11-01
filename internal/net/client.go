package net

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	timeout    time.Duration
}

func NewClient(timeoutSeconds float32) *Client {
	return &Client{
		httpClient: &http.Client{},
		timeout:    time.Duration(timeoutSeconds) * time.Second,
	}
}

func (c *Client) GetResponseBody(url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
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
