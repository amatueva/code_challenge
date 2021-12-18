package search

import (
	"context"
	"fmt"

	"main.go/internal/loader"
	"main.go/internal/types"
)

func Search(ctx context.Context, query types.Query) {
	switch query.Dataset {
	case "users":
		users := loader.LoadUsers(ctx)
		for _, u := range users {
			u.FindOne(query.Field, query.Value)
		}
	case "tickets":
		tickets := loader.LoadTickets(ctx)
		for _, t := range tickets {
			t.FindOne(query.Field, query.Value)
		}
	default:
		fmt.Println("something is wrong")
	}
}
