package search

import (
	"context"

	"main.go/internal/loader"
	"main.go/internal/types"
)

func Search(ctx context.Context, query types.Query) {
	records := loader.LoadData(ctx, query)
	records.FindOne(query.Field, query.Value)
}
