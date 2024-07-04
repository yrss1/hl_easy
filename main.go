package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"sync"
)

type Request struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type Response struct {
	ID      string            `json:"id"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Length  int               `json:"length"`
}

var requestStore sync.Map

func main() {
	http.HandleFunc("/proxy", handleProxy)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleProxy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	proxyReq, err := http.NewRequest(req.Method, req.URL, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	for key, value := range req.Headers {
		proxyReq.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	respHeaders := make(map[string]string)
	for key, values := range resp.Header {
		respHeaders[key] = values[0]
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	responseID := uuid.New().String()
	response := Response{
		ID:      responseID,
		Status:  resp.StatusCode,
		Headers: respHeaders,
		Length:  len(body),
	}

	requestStore.Store(responseID, struct {
		Request  Request
		Response Response
	}{
		Request:  req,
		Response: response,
	})

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
