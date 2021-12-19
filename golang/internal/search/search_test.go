package search

import (
	"context"
	"testing"

	"main.go/internal/types"
)

func TestSearch(t *testing.T) {
	validQuery := types.Query{
		Dataset: "users",
		Field:   "Id",
		Value:   "4",
	}
	ctx := context.Background()

	tests := []struct {
		name  string
		query types.Query
	}{
		{
			name:  "successful search",
			query: validQuery,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			Search(ctx, validQuery)
		})
	}
}
