package main

import (
	"context"
	"fmt"
	"log"
	"os"

	cohere "github.com/cohere-ai/cohere-go/v2"
	client "github.com/cohere-ai/cohere-go/v2/client"
)

func main() {
	apiKey := os.Getenv("COHERE_API_KEY")
	co := client.NewClient(client.WithToken(apiKey))

	resp, err := co.Chat(
		context.TODO(),
		&cohere.ChatRequest{
			Message: "What year was I born?",
		},
	)
	if err != nil {
		log.Fatalf("Error receiving response from ChatRequest %v", err)
	}

	fmt.Println(resp)

}
