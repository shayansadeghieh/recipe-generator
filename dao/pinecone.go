package dao

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/pinecone-io/go-pinecone/pinecone"
)

func CreateIndex(indexName string, pc *pinecone.Client, ctx context.Context) error {
	_, err := pc.CreateServerlessIndex(ctx, &pinecone.CreateServerlessIndexRequest{
		Name:      indexName,
		Dimension: 1024,
		Metric:    pinecone.Cosine,
		Cloud:     pinecone.Aws,
		Region:    "us-east-1",
	})

	if err != nil {
		return fmt.Errorf("error creating serverless index %v", err)
	}
	return nil
}

func CheckIfIndexExists(indexes []*pinecone.Index, indexName string, ctx context.Context) bool {
	for _, existingIndex := range indexes {
		if existingIndex.Name == indexName {
			return true
		}
	}

	return false
}

type pineconeQuery struct {
	Vector          []float64 `json:"vector"`
	IncludeMetadata bool      `json:"includeMetadata"`
	TopK            int       `json:"topK"`
}

type PineconeResp struct {
	Matches []pineconeMatches `json:"matches"`
}

type pineconeMatches struct {
	Id       string         `json:"id"`
	Score    float32        `json:"score"`
	Metadata recipeMetadata `json:"metadata"`
}

type recipeMetadata struct {
	RecipeIngredients  string `json:"recipeIngredients"`
	RecipeInstructions string `json:"recipeInstructions"`
	RecipeName         string `json:"recipeName"`
}

func QueryVectorDB(vector []float64) (PineconeResp, error) {
	apiKey := os.Getenv("PINECONE_API_KEY")
	url := fmt.Sprintf("https://%s/query", os.Getenv("PINECONE_INDEX_HOST"))
	vectorQuery := pineconeQuery{
		Vector:          vector,
		IncludeMetadata: true,
		TopK:            1,
	}

	// Marshal the query object to JSON
	jsonVector, err := json.Marshal(vectorQuery)
	if err != nil {
		return PineconeResp{}, fmt.Errorf("error marshalling query to JSON: %v", err)
	}

	// Init an HTTP POST request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonVector))
	if err != nil {
		return PineconeResp{}, fmt.Errorf("error creating HTTP request: %v", err)
	}

	// Set headers
	req.Header.Set("Api-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return PineconeResp{}, fmt.Errorf("error sending HTTP request: %v", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PineconeResp{}, fmt.Errorf("error reading vector query response body: %v", err)
	}

	// Decode the body into PineconeResp
	var pineconeResp PineconeResp
	err = json.Unmarshal(body, &pineconeResp)
	if err != nil {
		// Handle error decoding JSON
		return PineconeResp{}, fmt.Errorf("error unmarshaling pinecone response body %v", err)
	}

	return pineconeResp, nil

}
