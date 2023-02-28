package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	basePath = "http://localhost:8080/v1"
)

func makeRequest(method string, url string, body interface{}) (*http.Request, error) {
	var reader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal request body: %v", err)
		}
		reader = bytes.NewReader(b)
	}
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return req, fmt.Errorf("new request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func doRequest(req *http.Request, response interface{}) error {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("do http.request: %v", err)
	}
	defer res.Body.Close()
	if response != nil {
		br, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("read body: %v", err)
		}
		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("status not 200: %s", br)
		}
		if err := json.Unmarshal(br, response); err != nil {
			return fmt.Errorf("unmarshal response: %v", err)
		}
	}
	return nil
}
