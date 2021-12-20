package search

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"main.go/internal/types"
)

func TestSearch(t *testing.T) {
	validUserQuery := types.Query{
		Dataset: "users",
		Field:   "Id",
		Value:   "4",
	}

	validUserQueryNoResult := types.Query{
		Dataset: "users",
		Field:   "Id",
		Value:   "499999",
	}

	validTicketQuery := types.Query{
		Dataset: "tickets",
		Field:   "Type",
		Value:   "incident",
	}

	ctx := context.Background()

	tests := []struct {
		name       string
		query      types.Query
		recordType string
	}{
		{
			name:       "successful search users",
			query:      validUserQuery,
			recordType: "users",
		},
		{
			name:       "empty search results users",
			query:      validUserQueryNoResult,
			recordType: "users",
		},
		{
			name:       "successful search tickets",
			query:      validTicketQuery,
			recordType: "tickets",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			records := Search(ctx, tc.query)

			if tc.recordType == "users" {
				assert.IsType(t, types.UserRecords{}, records)
			} else if tc.recordType == "tickets" {
				assert.IsType(t, types.TicketRecords{}, records)
			} else {
				t.Errorf("test failed due to records unknown type - %T", records)
			}
		})
	}
}
