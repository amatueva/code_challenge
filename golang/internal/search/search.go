package search

import (
	"context"

	"main.go/internal/loader"
	"main.go/internal/types"
)

func Search(ctx context.Context, query types.Query) {
	records := loader.LoadData(ctx, query)

	for _, record := range records {
		attr := record.Attributes()
		val, ok := attr[query.Field]
		if val == query.Value && ok {
			record.Show()
		}
	}
}
