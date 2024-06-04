package main

import (
	"context"
	"log"
	"net/http"
	"os"

	coClient "github.com/cohere-ai/cohere-go/v2/client"
	pinecone "github.com/pinecone-io/go-pinecone/pinecone"
	dao "github.com/shayansadeghieh/recipe-generator/dao"
	handler "github.com/shayansadeghieh/recipe-generator/handler"
)

func main() {

	// Init pinecone vector database client
	pc, err := pinecone.NewClient(pinecone.NewClientParams{
		ApiKey: os.Getenv("PINECONE_API_KEY"),
	})
	if err != nil {
		log.Fatalf("Error init pinecone client %v", err)
	}

	ctx := context.Background()

	indexName := "recipes"
	indexExists, err := dao.CheckIfIndexExists(indexName, pc, ctx)
	if err != nil {
		log.Fatalf("Error checking if index exists %v", err)
	}

	if !indexExists {
		err = dao.CreateIndex(indexName, pc, ctx)
		if err != nil {
			log.Fatalf("Error creating index %v", err)
		}
	}

	// Init cohere client
	apiKey := os.Getenv("COHERE_API_KEY")
	co := coClient.NewClient(coClient.WithToken(apiKey))

	http.HandleFunc("/chatRequest", func(w http.ResponseWriter, req *http.Request) {
		handler.ChatRequest(w, req, co) // Pass the cohere client to the handler function
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
