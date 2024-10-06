// internal/cohere/client.go

package cohere

import (
	"context"
	"os"

	cohere "github.com/cohere-ai/cohere-go/v2"
	client "github.com/cohere-ai/cohere-go/v2/client"
)

var Client = client.NewClient(client.WithToken(os.Getenv("COHERE_API_KEY")))

func Chat(ctx context.Context, message string) (string, error) {

	resp, err := Client.Chat(ctx, &cohere.ChatRequest{
		Message: message,
	})
	if err != nil {
		return "", err
	}

	return resp.Text, nil
}
