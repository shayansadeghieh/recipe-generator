package dao

import (
	"context"
	"fmt"

	cohere "github.com/cohere-ai/cohere-go/v2"
	coClient "github.com/cohere-ai/cohere-go/v2/client"
)

func GenerateEmbeddings(text string, co *coClient.Client, context context.Context) ([][]float64, error) {

	resp, err := co.Embed(
		context,
		&cohere.EmbedRequest{
			Texts:     []string{"hello", "goodbye"},
			Model:     cohere.String("embed-english-v3.0"),
			InputType: cohere.EmbedInputTypeSearchDocument.Ptr(),
		},
	)

	if err != nil {
		return [][]float64{}, fmt.Errorf("error generating embeddings %v", err)
	}

	return resp.EmbeddingsFloats.Embeddings, nil
}
