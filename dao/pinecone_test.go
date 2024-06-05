package dao

import (
	"context"
	"reflect"
	"testing"

	"github.com/pinecone-io/go-pinecone/pinecone"
)

func TestCheckIfIndexExists(t *testing.T) {
	testCases := []struct {
		name      string
		indexes   []*pinecone.Index
		indexName string
		ctx       context.Context
		expOutput bool
	}{
		{
			name: "Index does not exist",
			indexes: []*pinecone.Index{
				{
					Name:      "foo",
					Dimension: 128,
				},
			},
			indexName: "bar",
			ctx:       context.Background(),
			expOutput: false,
		},
		{
			name: "Index exists",
			indexes: []*pinecone.Index{
				{
					Name:      "foo",
					Dimension: 128,
				},
			},
			indexName: "foo",
			ctx:       context.Background(),
			expOutput: true,
		},
	}

	// Iterate over test cases and run them
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Run the function being tested
			indexExists := CheckIfIndexExists(tc.indexes, tc.indexName, tc.ctx)

			if !reflect.DeepEqual(indexExists, tc.expOutput) {
				t.Errorf("Error checking if index exists: Got %v, but expected %v", indexExists, tc.expOutput)
			}

		})
	}
}
