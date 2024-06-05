package dao

import (
	"context"
	"fmt"

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

func CheckIfIndexExists(indexName string, pc *pinecone.Client, ctx context.Context) (bool, error) {
	indexes, err := pc.ListIndexes(ctx)
	if err != nil {
		return false, fmt.Errorf("error listing indexes %v", err)
	}

	for _, existingIndex := range indexes {
		if existingIndex.Name == "recipes" {
			return true, nil
		}
	}

	return false, nil
}
