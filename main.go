package main

import (
	"log"
	"net/http"
	"os"

	coClient "github.com/cohere-ai/cohere-go/v2/client"
	handler "github.com/shayansadeghieh/recipe-generator/handler"
)

func main() {

	apiKey := os.Getenv("COHERE_API_KEY")
	co := coClient.NewClient(coClient.WithToken(apiKey))

	http.HandleFunc("/chatRequest", func(w http.ResponseWriter, req *http.Request) {
		handler.ChatRequest(w, req, co) // Pass the cohere client to the handler function
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
