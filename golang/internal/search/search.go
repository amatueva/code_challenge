package search

import (
	"context"

	"main.go/internal/loader"
	"main.go/internal/types"
)

func Search(ctx context.Context, query types.Query) types.Record {
	records := loader.LoadData(ctx, query)
	return records.FindOne(query.Field, query.Value)
}
