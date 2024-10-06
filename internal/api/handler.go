// internal/api/handler.go
package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	cohere "github.com/cohere-ai/cohere-go/v2"
	cohereclient "github.com/cohere-ai/cohere-go/v2/client"
)

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Response string `json:"response"`
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	apiKey := os.Getenv("COHERE_API_KEY")
	if apiKey == "" {
		http.Error(w, "API key not found", http.StatusInternalServerError)
		return
	}

	co := cohereclient.NewClient(cohereclient.WithToken(apiKey))

	resp, err := co.Chat(context.TODO(), &cohere.ChatRequest{
		Message: req.Message,
	})

	if err != nil {
		log.Printf("Error invoking Cohere API: %v\n", err)
		http.Error(w, "Failed to communicate with AI model", http.StatusInternalServerError)
		return
	}

	response := ChatResponse{
		Response: resp.Text,
	}
	json.NewEncoder(w).Encode(response)
}
