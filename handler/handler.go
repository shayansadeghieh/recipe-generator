package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	cohere "github.com/cohere-ai/cohere-go/v2"
	coClient "github.com/cohere-ai/cohere-go/v2/client"
	"github.com/shayansadeghieh/recipe-generator/dao"
	wire "github.com/shayansadeghieh/recipe-generator/wire"
)

var prompt string = "You are a top-tier chef specializing in recipes geared towards diets, allergies, illnesses and conditions. A friend has asked you %s. Use the recipe name (%s), recipe ingredients (%s) and recipe instructions (%s) to provide the user with a recommendation for a nutritional meal. Make sure to include all of the amounts in the ingredients and instructions. Please be brief. "

func ChatRequest(w http.ResponseWriter, req *http.Request, co *coClient.Client, context context.Context) {
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

	// Generate embeddings for the message
	embeddings, err := dao.GenerateEmbeddings(message.Message, co, context)
	if err != nil {
		log.Fatal(err)
	}

	pineconeResp, err := dao.QueryVectorDB(embeddings[0])
	if err != nil {
		log.Fatal(err)
	}

	resp, err := co.Chat(
		context,
		&cohere.ChatRequest{
			Message: fmt.Sprintf(prompt, message.Message, pineconeResp.Matches[0].Metadata.RecipeName, pineconeResp.Matches[0].Metadata.RecipeIngredients, pineconeResp.Matches[0].Metadata.RecipeInstructions),
		},
	)
	if err != nil {
		log.Fatalf("error receiving response from ChatRequest %v", err)
	}
	if len(resp.Text) == 0 {
		log.Fatal("response text from ChatRequest is empty")
	}

	// Directly write the text to the response writer
	w.WriteHeader(http.StatusOK) // Set status code to 200
	_, err = w.Write([]byte(resp.Text))
	if err != nil {
		log.Fatalf("error writing response to user %v", err)
	}

}
