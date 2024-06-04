package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	cohere "github.com/cohere-ai/cohere-go/v2"
	client "github.com/cohere-ai/cohere-go/v2/client"
)

func ChatRequest(w http.ResponseWriter, req *http.Request, co *client.Client) {
	if req.Body == nil {
		log.Fatal("request body is empty")
	}

	defer req.Body.Close() // Close the body after reading

	// Read body into a byte array
	body, err := io.ReadAll(req.Body)
	if err != nil {
		// Handle error reading body
		log.Fatalf("error reading request body %v", err)
	}

	// Decode the body into message var
	var message map[string]string
	err = json.Unmarshal(body, &message)
	if err != nil {
		// Handle error decoding JSON
		log.Fatalf("error unmarshaling request body %v", err)
	}

	messageStr := message["message"] // Extract the message from the decoded JSON

	resp, err := co.Chat(
		context.TODO(),
		&cohere.ChatRequest{
			Message: messageStr,
		},
	)
	if err != nil {
		log.Fatalf("error receiving response from ChatRequest %v", err)
	}

	// Directly write the text to the response writer
	w.WriteHeader(http.StatusOK) // Set status code to 200
	_, err = w.Write([]byte(resp.Text))
	if err != nil {
		log.Fatalf("error writing response to user %v", err)
	}

}
