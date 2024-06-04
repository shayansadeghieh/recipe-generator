package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	cohere "github.com/cohere-ai/cohere-go/v2"
	coClient "github.com/cohere-ai/cohere-go/v2/client"
	wire "github.com/shayansadeghieh/recipe-generator/wire"
)

func ChatRequest(w http.ResponseWriter, req *http.Request, co *coClient.Client) {
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
	var message wire.ChatRequest
	err = json.Unmarshal(body, &message)
	if err != nil {
		// Handle error decoding JSON
		log.Fatalf("error unmarshaling request body %v", err)
	}

	resp, err := co.Chat(
		context.TODO(),
		&cohere.ChatRequest{
			Message: message.Message,
		},
	)
	if err != nil {
		log.Fatalf("error receiving response from ChatRequest %v", err)
	}
	if len(resp.Text) == 0 {
		log.Fatalf("response text from ChatRequest is empty")
	}

	// Directly write the text to the response writer
	w.WriteHeader(http.StatusOK) // Set status code to 200
	_, err = w.Write([]byte(resp.Text))
	if err != nil {
		log.Fatalf("error writing response to user %v", err)
	}

}
