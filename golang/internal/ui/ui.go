package ui

import (
	"context"
	"encoding/json"
	"fmt"

	"main.go/internal/types"
	"main.go/internal/validation"
)

func PromptUser(ctx context.Context) (types.Query, error) {
	var dataset string
	var field string
	var value interface{}

	dataset = promptForDataset(ctx)

	if dataset == "users" {
		field = promptForUserField(ctx)
	} else {
		field = promptForTicketField(ctx)
	}

	value = promptForSearchValue(ctx)

	return types.Query{Dataset: dataset, Field: field, Value: value}, nil
}

func promptForDataset(ctx context.Context) string {
	for {
		go func() {
			<-ctx.Done()
			fmt.Println("context canceled. Exiting")
		}()
		var dataset string
		fmt.Println("Are you searching for tickets or users?")
		fmt.Scanln(&dataset)
		if validation.ValidateDataset(dataset) {
			return dataset
		}
	}
}

func promptForUserField(ctx context.Context) string {
	for {
		go func() {
			<-ctx.Done()
			fmt.Println("context canceled. Exiting")
		}()
		var field string
		fmt.Println("Type user property that you're searching")
		fmt.Scanln(&field)
		if validation.ValidateUserField(field) {
			return field
		}
	}
}

func promptForTicketField(ctx context.Context) string {
	for {
		go func() {
			<-ctx.Done()
			fmt.Println("context canceled. Exiting")
		}()
		var field string
		fmt.Println("Type ticket property that you're searching")
		fmt.Scanln(&field)
		if validation.ValidateTicketField(field) {
			return field
		}
	}
}

func promptForSearchValue(ctx context.Context) interface{} {
	for {
		go func() {
			<-ctx.Done()
			fmt.Println("context canceled. Exiting")
		}()
		var value string
		fmt.Println("Type value you're seacring for")
		fmt.Scanln(&value)
		if validation.ValidateSearchValue(value) {
			//parse value as json and return it
			var parsedValue interface{}
			err := json.Unmarshal([]byte(value), &parsedValue)
			if err == nil {
				return parsedValue
			}
		}
	}
}
